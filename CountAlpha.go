package main

import "fmt"

func CountAlpha(s string) int {
	nbr := 0
	for _, v := range s {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			nbr++
		}
	}
	return nbr

}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

/*Write a function CountAlpha() that takes a string as an argument and returns the number of alphabetic characters in the string.*/
