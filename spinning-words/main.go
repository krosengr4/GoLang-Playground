package main

import (
	"fmt"
	"strings"
)

/*
- Write a function that takes in a string of one or more words.
- Return the same string, but for every word with 5 or more letters, it is reversed.
- Ex: "Hello this is Rachel" -> "olleH this is lehcaR"
*/

func main() {
	str := "Stop spinning my words"
	result := spinWords(str)
	fmt.Println(result)
}

func spinWords(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		r := []rune(word)
		if len(r) >= 5 {
			for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
				r[i], r[j] = r[j], r[i]
			}
			words[i] = string(r)
		}
	}

	result := strings.Join(words, " ")
	return result
}
