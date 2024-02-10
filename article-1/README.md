## Try-Catch Behavior in Golang

In general the most of languages have a mechanism  to catch errors, Java, Python, JS etc. But when we talk about Golang thats are huge different the Golang keep the flow and deal with error as another return. However exist the fatal/pania error who got a different behavior, when is raised error execition is stoped.


I supposed that you are asking why do i want handle fatal error if they purpose is exactly stop the execution. I have asked myself either. Fist off i think you should not use that way to handle fatal error in golang except when is necessary. And i think that way is kind similiar with ther error handling of most language outside.

How golang handling with error in a regular situation?

Golang has different behavior from other languages, the error is seen as normal value and not an exception. This way the flow is quite more simple to understand in golang. I will provide two exemple below of script to read a file 


Python

```
path_file = "exemple.txt"
content = None
with open(path_file, "d") as f:
    content = f.read()


# With try catch
try:
   with open(path_file, "d") as f:
    content = f.read()
except FileNotFoundError:
    print("Error while try reading file")
```

Golang
```
package main

import (
	"fmt"
	"os"
)
const filePath = "exemple.txt"
func main() {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error while try reading file")
	}
	fmt.Println(string(content))
}

```


Ok but if i want implement a similiar try catch behavion in Golang? Case we ve got a fatal error in golang, that error is a unespected raised error who stop the execution. In other words, has same behavior of any another error in other languages. To this exemple i will get the string character by index. When i got an index not exist will raised fatal error up  



Python
```
word = "matheus"
index = 30
try:
    char = word[30]
    print(f"this is {char}")
except:
    print(f"Index {str(index)} was not exist in  the string '{word}'")
```

Golang
```
package main

import "fmt"

func getCharIndex(str string, idx int) (char rune, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Index %d was not exist in  the string '%s'", idx, str)
		}
	}()

	char = rune(str[idx])
	return char, nil
}

func main() {
	word := "Matheus"
	char, err := getCharIndex(word, 30)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("this is '%c' char", char)
	}

}

```

Both cases above has same behavior, try use a index off of string raised a error and catch the error without stop the program. But while in Python is very used see code struct like that and in golang in golang is has rarely seem.

Basically to create the try catch behavior utilizing recover function who will catch panic case exist, notice that funciton required be with defer. So the flow in getCharIndex function, the function run till raised error up then will be execute defer where has the recover funcion to catch the error

If your code depends try-catch behavior your application get slow because the overhead of exception handling has more cost. Assure use try-catch just in not excpeted situations, not for regular control control flow.

Now get back to our work around try-catch suppose implemented in especial cases, because case your golang code is using this behavior when you can avoid the error you are implemeting wrong way. Now i will re-write my go code using best pratices

```
package main

import "fmt"

func getCharIndex(str string, idx int) (char rune, err error) {
	if len(str) <= idx {
		err = fmt.Errorf("index out of range: %d", idx)
	} else {
		char = rune(str[idx])
	}
	return char, err
}

func main() {
	str := "Matheus"
	char, err := getCharIndex(str, 30)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("this is '%c' char", char)
	}

}

```





