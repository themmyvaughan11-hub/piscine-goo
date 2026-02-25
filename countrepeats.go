package main

import "fmt"

func CountRepeats(s string) int {
	if len(s) == 0 {
		return 0
	}

	nbr := 0

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			nbr++
		}
	}
	return nbr
}

func main() {
	fmt.Println(CountRepeats("aaabbcdddd"))
	fmt.Println(CountRepeats("abc"))
	fmt.Println(CountRepeats("aabbcc"))
}

/*Write a function countrepeatts
that takes a string as an argumant and returns the  number of consectutive  repeated characters in the string
only consecutive repeated characters must be counted
  If a character appers more than once in a row , each additional repetition increase the count

  The function must return the total number of repeated characters found */
