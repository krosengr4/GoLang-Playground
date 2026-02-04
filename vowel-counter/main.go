package main

import (
	"fmt"
	"strings"
)

/*
- Given a list of words
- For each word, print how many vowels there are for that word
- Print out how many total vowels there are in the entire list
- Vowels: a, e, i, o, u AND y
*/

func main() {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("\t-----VOWEL COUNTER-----")

	words := []string{"telescope", "whisper", "umbrella", "cinnamon", "glacier", "rhythm", "lantern"}

	totalVowels := vowelCount(words)

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Total vowels in list: ", totalVowels)
	fmt.Println(strings.Repeat("-", 50))
}

func vowelCount(words []string) int {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("WORD LIST:\n", words)
	fmt.Println(strings.Repeat("-", 50))

	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'y'}
	totalVowelCount := 0

	// Loop through each word
	for _, word := range words {
		vowelCount := 0
		// Loop through each char in word
		for _, char := range word {
			// Loop through each vowel and check for vowels
			for _, vowel := range vowels {
				if char == vowel {
					vowelCount++
					totalVowelCount++
				}
			}
		}

		fmt.Printf("%s: %d vowels\n", word, vowelCount)
	}

	return totalVowelCount	
}
