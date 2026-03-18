package main

import (
	"fmt"
	"strings"
)

/*
- Given a string made up of words and spaces, return the length of the last word
- A word is a maximal substring consisting of non-space characters only.
*/

func main() {
	fmt.Println("WHAT IS THE LENGTH OF THE LAST WORD?!")

	str := "Hello my name is Gina Mckgirbby hello   "
	length := lengthOfLastWord(str)

	fmt.Println("Last Word Length:", length)
}

func lengthOfLastWord(s string) int {
	// splits on any whitespace and trims leading/trailing spaces automatically
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0
	}

	result := len(words[len(words)-1])
	return result
}
