package main

import (
	"fmt"
	"strings"
)

// Given a list of random integers, sort the list from smallest to biggest
// Use bubble sort
// Print the sorted list

func main() {
	fmt.Println(strings.Repeat("-", 55))
	fmt.Println("BUBBLE SORT")
	fmt.Println(strings.Repeat("-", 55))

	nums := []int{5, 3, 19, 12, 3}
	
	result := bubbleSort(nums)
	for i := 0; i < len(result); i++ {
		fmt.Printf("%d: %d\n", i + 1, result[i])
	}

}

func bubbleSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums) - 1; j++ {
			if nums[j] > nums[j + 1] {
				nums[j], nums[j + 1] = nums[j + 1], nums[j]
			}
		}
	}

	return nums
}
