package main

import "fmt"

func getCharacter(str string, index int) (char rune, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("attempted to access index %d out of range", index)
		}
	}()

	char = rune(str[index])
	return char, nil

}
func main() {
	aa := fmt.Sprintf("%c", 46)
	fmt.Println(aa)
	char, err := getCharacter("Hello World!", 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The  character at index 50 is '%c'\n", char)
	}

}
