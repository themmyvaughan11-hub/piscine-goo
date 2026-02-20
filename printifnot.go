package main

import "fmt"

func PrintIfNot(str string) string {
	if len(str) < 3 {
		return "G\n"
	}
	return "Invalid input\n"
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}

/*Write a function that takes a string and returns:

"G\n" if the string's length is less than 3 (including empty string).

"Invalid Input\n" otherwise.*/
