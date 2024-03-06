## Combining Factory and Functional Options Pattern in Golang

Theses patterns are widely used in Golang and, when combined, can be extremely powerful. Constructing objects in a flexible manner allows for defaults and the ability to customize the object being created through options.

In this blog post, we will explore:

1 - Factory Pattern 
2 - Functional Options Pattern
3 - How to Write Generic Helper Functions
4 - Implementing Both Design Patterns Together.

## Short Factory Pattern introduction

If you are programming code in Golang probably you seen function with "New" prefix, like "NewUser" to create a User. If have saw that, then you know how Factory Pattern works, even if you are not familiar with the term 'Factory Pattern. 

### In this blog post, we will explore:
Back to initial concept of Factory Pattern , it is a method that create a concrete Object that implements a common interface or a specific struct. This pattern is particulary useful in maintaining a clean and organized codebase, promoting the principle of DRY(Don't Repat Yourself) by centralizing the creation logic.

In a exemple case where we have a struct called Server, imagine that every time you use the Server struct, it's necessary to pass its fields. But if the field values are always the same, why pass them into a server every time to use the Server? To enhance this process, the factory pattern is used. This way, we can just create a method called NewLocalServer, wich will return the server.



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


The snippet code above represent a simple use case of Factory Pattern. Now we have a concrete Server, i don't need worry about set up the server.


However, an issue arises, Lest's suppose that now i want create a LocalHost server but with a different timeout or port value. In this case just the factory pattern wont be able to ACCOMMODATE IT. I could create multiples method factory with diferents timeoutS, but is not the puporse of factory method. Then now to fix it the Function Option Pattern come into play.




## What is Function Option Pattern

The functional options pettern is a design pattern thats allows flexibiliy when initialization objects in Golang. They can used as required or optional arguments depends on the goals. 

Let's code show use cases

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

The code snipped above has used Function Option Pattern, now it's possible personalise through Generic Helper Functions WithTimeOut or WithPort the object returned by the factory method. Now lets break implementation down in two part. How create a Generic Helper Function and how implement them into a factory method.


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

## Conclusion

















