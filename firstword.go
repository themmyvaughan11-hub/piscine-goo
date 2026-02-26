package main

import "fmt"

func FirstWord(s string) string {
	if len(s) == 0 {
		return s + "\n"
	}
	for i, v := range s {
		if v == ' ' {
			return s[:i] + "\n"

		}
	}
	return s + "\n"
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

/*Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').

A word is a sequence of characters delimited by spaces or by the start/end of the argument.*/
