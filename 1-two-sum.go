package leetcode_golang

/*
1. Two Sum

https://leetcode.com/problems/two-sum/description/
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func twoSum(nums []int, target int) []int {
	intMap := make(map[int]int)

	for currIdx, currVal := range nums {
		diff := target - currVal
		if value, ok := intMap[diff]; ok {
			return []int{value, currIdx}
		} else {
			intMap[currVal] = currIdx
		}
	}

	return []int{}
}
