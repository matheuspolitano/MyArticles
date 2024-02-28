## Unlocking the Power of Functional Options Pattern in Go

If you're new to Golang, understanding patterns is particularly crucial for navigating popular libraries, SDKs and writing code that's flexible, scalable, customizable, and maintainable. One pattern that stands out, especially in the Go ecosystem, is the Functional Options Pattern. This pattern is not commonly found in many other programming languages, so even if you're an experienced engineer, you might not be familiar with it.

In this blog post, we'll dive into

1. What Problems Does the Functional Options Pattern Solve?
2. Functional Options Pattern Concept
3. How to Write Generic Helper Functions
4. Implementing into an factory method.

## What Problems Does the Functional Options Pattern Solve?

In Go, we have the flexibility to create new types based on structs, allowing us to specify only the fields we need. For example

```
package main

import (
	"log"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
}

func (s *Server) Run() {
	log.Printf("Server running %s:%d", s.host, s.port)

}
func (s *Server) Stop() {
	log.Printf("Server has stopped %s:%d", s.host, s.port)
}

func main() {
	localHostServer := &Server{
		host:    "127.0.0.1",
		port:    8080,
		timeout: 3 * time.Second,
	}
	localHostServer.Run()
}

```

### Creating default values

Now, consider a scenario where we frequently create a Server instance with the same attribute values. Unfortunately, we can't directly encode these default values into the type itself. This is where the Factory Method comes into play. In essence, the Factory Method is a design pattern that involves a method dedicated to creating and returning an instance of an object, pre-populated with default values. For example

```

func NewLocalHost() *Server {
	return &Server{
		host:    "127.0.0.1",
		port:    8080,
		timeout: 3 * time.Second,
	}

}

func main() {
	localHostServer := NewLocalHost()
	localHostServer.Run()
}



```

Notice that we now have a method to generate a Server object with predefined values, which is incredibly useful given that these values are common across my application. Thus, we've introduced a function, NewLocalHost, which provides us with an object that's ready to use

### Now How Can We Modify Default Values?

The dilemma often boils down to 'can't' versus 'shouldn't'. Despite the constraints, a practical workaround involves passing arguments to a Factory Method


**It's advisable to steer clear of implementing it in that manner.**
```
// NewLocalHost creates a new Server instance with optional port and timeout parameters.
// If port or timeout are not provided (nil), default values are used.
func NewLocalHost(port interface{}, timeout interface{}) *Server {
	defaultPort := 8080
	defaultTimeout := 3 * time.Second

	// Check and set port if provided
	actualPort := defaultPort
	if p, ok := port.(int); ok {
		actualPort = p
	}

	// Check and set timeout if provided
	actualTimeout := defaultTimeout
	if t, ok := timeout.(time.Duration); ok {
		actualTimeout = t
	}

	return &Server{
		host:    "127.0.0.1",
		port:    actualPort,
		timeout: actualTimeout,
	}
}

func main() {
	// Example usage of NewLocalHost without parameters, using default values
	localHostServer := NewLocalHost(9090, nil)
	localHostServer.Run()

	// After some operations, stop the server
	// localHostServer.Stop()
}

```

Directly passing field values to a factory method, rather than leveraging the Functional Options pattern, can lead to several challenges, especially as your codebase expands and evolves. Here are some potential issues associated with this approach:



`Long Parameter List`: As the number of server parameters increases, the factory method's signature becomes unwieldy, making it difficult to manage and use.

`Limited Flexibility for Defaults`: It becomes cumbersome to manage default values. Direct parameter passing forces callers to specify all values explicitly, including those that should default, or necessitates multiple constructors for different scenarios, cluttering the code and reducing maintainability.

`Compromised Readability`: Calling a function with multiple parameters, especially of the same type, obscures the purpose of each parameter without referencing the function definition, detracting from code readability and increasing the likelihood of errors.

`Reduced Encapsulation and Flexibility`: Exposing the internal structure and implementation details through direct parameter passing limits the flexibility to modify the internal implementation without affecting the interface.

`Inconsistent Object State`: The absence of a structured mechanism to enforce necessary field settings or validate configurations can lead to objects being in an inconsistent or invalid state, posing significant risks to application stability.



## How Function Options Pattern solve all of those issues?

This pattern involves functions that you can aggregate with a Factory Method. Here's an example:

```
func main() {
	localHostServer, err := NewLocalHost(WithPort(9090))
	if err != nil {
		log.Fatal(err)
	}
	localHostServer.Run()
}
```


The code snippet mentioned leverages the Functional Options Pattern, allowing for customization of values through Generic Helper Functions like WithTimeOut or WithPort.

It's important to note that we've only showcased the main function using the Options Pattern, not the entire implementation. The goal here is to highlight how straightforward it is to use these methods instead of passing values directly. Some advantages of using this pattern include:

`Enhanced Readability`: By using named options, the code becomes self-documenting. You can easily understand what configuration is being applied without digging into the details of each function.


`Flexibility in Configuration`: You can provide defaults and allow the consumer of your code to override them as needed without changing the function signature.


`Scalability`: Adding new options is simple and doesn't break existing code. You can introduce additional functionalities without affecting the consumers of your code.


`Improved Code Maintenance`: Since the options are functions, they encapsulate the logic for setting up the object, reducing the complexity of the initialization code and making it easier to maintain.


`Decoupled Code`: This pattern helps in keeping your code decoupled and easy to test, as you can mock these options in your unit tests.



In summary, the Functional Options Pattern offers a robust, maintainable, and flexible approach to configuring objects in Go, making it a valuable technique for building scalable and clean applications

However, what remains to be addressed is the comprehensive implementation details, specifically how to create a `Generic Helper Function` and how to seamlessly integrate these functions into a Factory Method.


### How create a Generic Helper 

An important aspect to note about Generic Helper Functions is that they do not directly modify any values. Instead, these functions return another function that takes the object itself as an argument and may return an error. This design ensures that any changes to the object's state occur in a subsequent step, maintaining a clear separation of concerns.

Let's proceed by defining the type returned by a Generic Helper Function:

```
type OptionsServerFunc func(c *Server) error
```

Now that we've established our custom type, let's refine the WithPort Generic Helper Function, incorporating best practices and improvements. Adhering to these guidelines enhances clarity and consistency:

- Prefix the function name with With followed by the name of the field being modified.
- Ensure the function receives only one argument, which should be of the same type as the field it is intended to modify.
- The function must return another function that performs the assignment operation.

To illustrate how error handling can be integrated, we'll add validation to ensure the port is within the range of 5000 to 9999.

Let's dive into the coding process:

```
func WithPort(port int) OptionsServerFunc {
    return func(s *Server) error {
        // Check if the port is within the valid range.
        if port >= 5000 && port < 10000 {
            s.port = port
            return nil
        }
        // Return an error if the port is out of the valid range, using formatted error string for clarity.
        return fmt.Errorf("port %d is out of the valid range (5000-9999)", port)
    }
}
```

The code snippet presents a Generic Helper Function that's prepared for use. However, even if we pass it as an argument, it won't function within NewLocalHost without an existing mechanism to handle it


### How add mechanism to handle with Generic Helper Fuction into a factory method?


To handle multiple Generic Helpers that are optional, it's essential to iterate through each one, applying them to the object and verifying whether they return an error.

Firstly, to enable a function to accept multiple arguments, we utilize a feature known as 'variadic functions'. A variadic function can receive an indefinite number of arguments of the same type, facilitated by prefixing the parameter type with an ellipsis ('...') in the function's definition.

Here's an example of a variadic function:

```
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2))
    fmt.Println(sum(1, 2, 3))

    // You can also pass a slice of ints by using the ellipsis suffix
    numbers := []int{1, 2, 3, 4}
    fmt.Println(sum(numbers...))
}
```


Now that we understand how to utilize the variadic mechanism, we need to iterate over the list of Generic Helper Functions. This can be done using a range loop, where we can ignore the index and pass the newly created object as an argument to each function.

Let's proceed with coding this implementation:
```
func NewLocalHost(opts ...OptionsServerFunc) (*Server, error) {
	server := &Server{
		host:    "127.0.0.1",
		port:    8080,
		timeout: 3 * time.Second,
	}

	for _, opt := range opts {
		if err := opt(server); err != nil {
			return nil, err
		}
	}
	return server, nil

}
```

The code snippet features a loop within the NewLocalHost function that iterates over a slice of OptionsServerFunc. Each opt within the slice is a function accepting a pointer to a Server as its argument, and it returns an error. As the loop progresses, each opt is invoked with server as its parameter. Should any opt return an error, the loop halts prematurely, and the NewLocalHost function returns both nil and the encountered error.

This design facilitates a modular and flexible configuration of the Server instance, allowing for a sequence of modifications or validations as defined by the OptionsServerFunc in the opts slice.

Now we have all of Functional Options Pattern implemented :)

```
type OptionsServerFunc func(c *Server) error

func WithTimeout(t time.Duration) OptionsServerFunc {
	return func(c *Server) error { c.timeout = t; return nil }
}

func WithPort(p int) OptionsServerFunc {
	return func(c *Server) error { c.port = p; return nil }
}
func NewLocalHost(opts ...OptionsServerFunc) (*Server, error) {
	server := &Server{
		host:    "127.0.0.1",
		port:    8080,
		timeout: 3 * time.Second,
	}

	for _, opt := range opts {
		if err := opt(server); err != nil {
			return nil, err
		}
	}
	return server, nil

}

func main() {
	localHostServer, err := NewLocalHost(WithTimeout(5*time.Second), WithPort(7000))
	if err != nil {
		log.Fatal(err)
	}
	localHostServer.Run()
}
```

## Conclusion

Throughout our exploration, we've delved into the core aspects of this pattern, shedding light on its widespread adoption in the Go programming landscape. It's my hope that this discussion has illuminated the subject for you. Remember, the path to mastering programming is paved with practice. Stay curious, keep coding, and let's continue to evolve our skills together