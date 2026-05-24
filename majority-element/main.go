package main

import "fmt"

/*
- Given an array of nums of size n
- Return the majority element
- Majority element is the one that appears more than n/2 times.
*/

func main() {
	arr := []int{4, 4, 1, 1, 1, 4, 4}
	result := majorityElement(arr)

	fmt.Println(result)
}

func majorityElement(arr []int) int {
	counts := make(map[int]int)

	for _, num := range arr {
		counts[num]++
	}

	threshold := len(arr) / 2
	for num, count := range counts {
		if count > threshold {
			return num
		}
	}

	return -1
}
