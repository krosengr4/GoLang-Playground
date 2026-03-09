package main

import "fmt"

/*
- Given an array of integers nums and an integer target, return indicies of the two numbers such that they add up to the target
- You may assume that each input has exactly one solution, and you may not use the same element twice
- Return the answer in any order
*/

func main() {
	nums := []int{4, 3, 8, 9, 8, 10}
	target := 16

	result := twoSum(nums, target)
	fmt.Println("Result:", result)
}

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}
	for i, num := range nums {
		complement := target - nums[i]
		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}

		seen[num] = i
	}
	return nil
}
