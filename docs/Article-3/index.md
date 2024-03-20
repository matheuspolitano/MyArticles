## Elevate Your Go Logging: Why Zerolog is Your Best Choice

Imagine your first code - likely, it was a simple  "Hello World" printed in a log. Logging is an insdipensable part of every application, whether its a vast enterprise systen ir a modest callege project. Even amidst complex services and metrics handling, logging remains a cornerstone. This bring us to Zerolog, a high-performance logging library for Go, celebrated for its speed and ease of implmentation

 
 Quick note this is not a common guideline approach where use hard term to explain simple things. I will show you how use zerolog of course, but its not just it. Zerolog use some interting patters and if you get it, you can maybe developer other things using thoses patterns. So take your sit that we gonna take off our article 


## Topics

- Why Zerolog is so fast?
- My first zerolog
- Zerolog 
- How implement in 



# Why Zerolog is so fast?

Zerolog is designed to avoid memory allocations as much as possible. This design priciple is crucial for achieving high performance in logging system. By minimizing allocation adn reduce pressure on the garbage collectior. leading to less GS pause times and improced overal application performance.

Implementation take advantage of modern CPU architetctures, using eddicient algorithms and data structure that reduce computational overhead and improve cache utilization



# My first zerolog



```
package main

// importing default log API from zerolog
import "github.com/rs/zerolog/log"

func main() {

	//API zerolog - Factory Mehod - Add fields - Emit Log 
	log.Debug().Str("Msg", "My first log").Str("Msg","Running main").Send()
}

```

i've included a snipped code right at the outset for you see firsthand how strightforward it is to uderstand Zerolog. It's incredibly intuitive, allowing you to code effectively even whithout extensive background knowledge

At first looking, we can see the log who is default zerolog API, and calling the method `Debug` who is a Factory Method. In essence, the **Factory Method** is a design pattern that involves a method dedicated to creating and returning an instance of an object, pre-populated with default values. For example in this case returning an Event with trace type return



```
package main

// importing default log API from zerolog
import "github.com/rs/zerolog/log"

func main() {

    \\Whithout Factory Method
    log.newEvent(DebugLevel, nil).Str("Msg", "My first log").Str("Msg","Running main").Send()
}   
```

Then as you can see the Factory help us create event with less code. 


 you can see series of method calls are chaining together. That is used pattern in Golang Called "Fluent Interface"






















