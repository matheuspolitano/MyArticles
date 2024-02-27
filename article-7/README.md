## Unlocking the Power of Functional Options Pattern in Go

If you're new to Golang, understanding patterns is particularly crucial for navigating popular libraries and writing code that's flexible, scalable, customizable, and maintainable. One pattern that stands out, especially in the Go ecosystem, is the Functional Options Pattern. This pattern is not commonly found in many other programming languages, so even if you're an experienced engineer, you might not be familiar with it.

In this blog post, we'll dive into

1 - What Problems Does the Functional Options Pattern Solve?
2 - Functional Options Pattern Concept
3 - How to Write Generic Helper Functions
4 - Implementing into an factory method.

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

Now imagine that the most of the time we gonna create a Server with the sames atributes values, we can't pass the default values through type. Then we gonna use Factory Method, in a short explanation is a method that will create an object with values. Example

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

Notice that now we have a way to generate a Serve object with value, thats very useful because i inteed that's values is very common to my application. So now we have a function `NewLocalHost` who give to us the object ready to use.

### Now How Can We Modify Default Values?

You can't or you shouldn't, that's the issue. We could work around and pass argument to a Factory Method. 


**Avoid implement that way**
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

Passing field values directly to a factory method instead of using the functional options pattern can introduce several issues, particularly as your codebase grows and evolves. Here are some of the problems you might encounter with this approach:



- `Long Parameter List`: As the number of server parameters grows, the factory method signature become unwildy
- `Limited Flexibility for Defaults`: Managing default values becomes cumbersome. With direct parameter passing, you either force callers to specify all values explicitly, including those that should often just be defaults, or you create multiple constructors for different scenarios, which leads to cluttered and less maintainable code.

- `Compromised Readability:` When a function is called with multiple parameters, especially if they are of the same type, it's hard to tell what each parameter represents without looking up the function definition. This makes the code less readable and more error-prone.

- `Reduced Encapsilation and Flexibility`: Directly passing parameters requires exposing the internal structure and implementation details of your objects. This can reduce the dlexibility to change the internal implementation.

- `Inconsistent Object State:` Without a clear mechanism to enforce the setting of necessary fields or validate the configuration, it's easy to end up with objects in an inconsistent or invalid state. 



## How Function Options Pattern solve all of those issues?

This pattern are functions that you can agregate to a Factory Method an Example

```
func main() {
	localHostServer, err := NewLocalHost(WithPort(9090))
	if err != nil {
		log.Fatal(err)
	}
	localHostServer.Run()
}
```


The code snipped above has used Function Option Pattern, now it's possible personalise the values through `Generic Helper Functions` WithTimeOut or WithPort. 


Notice that we just have showed the `main` using the Options Pattern not the over all implementation. I've wanted point out how easy is use this methods instead pass the value directly. Couples advantages of use that Pattern:

- `Flexible`: It allows for easy addition of new options without breaking existing code.

- `Scalable`: The pattern scales well with complex configurations and evolving software requirements.

- `Readable`: Code using functional options is often more readable than alternatives, making it easier to understand what options are being set.

- `Intuitive`: The pattern leverages Go's first-class functions and closures, making it intuitive for those familiar with these concepts.

- `Customizable`: Offers a high degree of customization, allowing developers to define options that can precisely control the behavior of their objects.

- `Maintainable`: The pattern promotes maintainability by keeping configuration logic centralized and decoupled from the object's core functionality.




But is missing the full implementation, `how create a Generic Helper  Function` and `how implement them into a Factory Method`?




### How create a Generic Helper 

A import point about Generic Helper Functions, they dont execute any change to value that process is going just on the next step. They return a function that receive as argument own object and return an error. 

Then let's create a type return by Generic Helper Function

```
type OptionsServerFunc func(c *Server) error
```

Now that we have own type let's recreate the `WithPort` Generic Helper with an improvement. Some patterns and 

-  Prefix name as `With` and field changed
-  Receive only one argument per function
-  Argument must has same type of the field who will receive
-  Must return a function with the assign operation


To show as error mechanism can be used we will add extra validation to check if the port range of 5000 to 9999,


Let's code

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

The code snipped has a Generic Helper Fuction ready to be used even thow if we use as argument won't work into NewLocalHost if mechanism to receive not exist.


### How add Generic Helper Fuction into a factory method?


We need deal with multiples Genreic Helper who are optional and then read one by one passing into them the object and check if return an error value. 

Fist to the function receive multiples arguments we must use a functionality is called "variadic functions". A variadic function can take an arbitrary number of argument of the same type. This achieved by using the ellipsis ('...') prefix before the parameter type in the function definition.


Variadic Function exemple 

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


Now we've already know how use the variadic mechanism we need iterate with the list of Generic Help Function. We can use a range to do it, ignore the index and pass the objected created as argument.


Let's code
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

















