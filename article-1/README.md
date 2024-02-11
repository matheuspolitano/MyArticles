# How use try-catch mechanism in Golang?

Unlike most programming languages, such as Java, Python, and JavaScript, which rely on a try-catch mechanism for error handling, Golang treats errors in a markedly different way. In Golang, errors are simply regarded as normal return values, integrating smoothly with the language's flow.

# Error Handling in Normal Situations in Golang
Golang's method for handling errors stands out from other languages. Instead of treating errors as exceptions, Golang considers them as regular values. This approach streamlines the control flow, making the behavior of programs more predictable and generally easier to understand.

Here, I will showcase two examples of scripts that illustrate how to read a file in Golang, demonstrating the conventional method for handling errors in this language.

Python
```
path_file = "exemple.txt"
content = None
try:
   with open(path_file, "r") as f:
    content = f.read()
    print(content)
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

## How handle with panic?

This approach contrasts sharply with Golang's more direct method of handling fatal errors or panics, which stops the execution immediately upon encountering such severe problems.

If one aims to implement something akin to a try-catch mechanism in Golang, particularly for handling fatal errors—those unexpected, uncommanded errors that cease the operation—it's crucial to understand that these errors are treated similarly to standard errors in other programming languages.

Why handle fatal errors if their intent is to halt execution?
This question has also occurred to me. Initially, it's vital to point out that managing fatal errors in Golang should only be considered when absolutely necessary. Nevertheless, devising a strategy to handle these errors could make Golang's error management more comparable to that found in other languages.

For instance, attempting to access a string character at a non-existent index would provoke a fatal error. I'll be showcasing two examples: one using Python to highlight the conventional approach adopted by numerous languages, and another using Go to exhibit its distinctive approach to addressing the same issue.



Python
```
word = "matheus"
index = 30
try:
    char = word[30]
    print(f"this is {char}")
except:
    print(f"Index {str(index)} was not exist in the string '{word}'")
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


 Both cases above display the same behavior, where an attempt to use an index off of a string results in an error, yet the error is caught without stopping the program. However, while this kind of code structure is quite common in Python, it's rarely seen in Golang.

Essentially, to implement try-catch behavior, the recover function is used to catch any panic that may occur, with the important note that this function needs to be used with defer. Therefore, in the flow of the getCharIndex function, the function runs until an error is raised, and then the defer statement executes, which includes the recover function to catch the error.

If your code relies on try-catch behavior, your application will slow down due to the additional overhead of exception handling. It's crucial to use try-catch only in unexpected situations, not as a regular method for controlling flow.

Now, let's return to our discussion on try-catch, which should be implemented only in special cases. If your Golang code uses this behavior when you could avoid the error, you're implementing it incorrectly. Now, I will rewrite my Go code using best practices.

Golang
```
package main

import "fmt"

func getCharIndex(str string, idx int) (char rune, err error) {
	if len(str) > idx {
		char = rune(str[idx])
	} else {
		err = fmt.Errorf("Index %d was not exist in  the string '%s'", idx, str)
	}
	return char, err
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

Reviewing the changes, it is observed that the 'err' and 'char' functionalities remain intact. The significant adjustment is the elimination of the 'recover' method, which previously acted as a mechanism for detecting errors.

 The reason for removing this method is based on the understanding that if errors can be anticipated and prevented, then relying on a method to catch errors after they occur is rendered unnecessary. Consequently, a check was introduced to ascertain the length of the string, ensuring that the index is within the permissible bounds of the string. If the index is less than the length of the string, it indicates that the index is valid and part of the string; otherwise, it results in an error. 
 
 This method represents a shift towards active error management, as opposed to the former approach where the system passively awaited the occurrence of an error. Adopting this proactive approach to error handling is in line with the best practices recommended by Go programming.


## Conclusion

So, in this article, I've shown how errors are handled in Go and explored the aspects of implementing a try-catch-like behavior as a workaround, along with their applicability to use cases. I hope this discussion has been interesting and useful for you.