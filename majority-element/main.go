package main

import (
	"fmt"
	"slices"
)

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
	slices.Sort(arr)

	arrLength := len(arr)

	return arr[arrLength/2]

}
