package leetcode_golang

import (
	"testing"
)

func Test_twoSum(t *testing.T) {

	type testData struct {
		nums     []int
		target   int
		solution []int
	}

	data := []testData{
		{nums: []int{2, 7, 11, 15}, target: 9, solution: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, solution: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, solution: []int{0, 1}},
	}

	for _, value := range data {
		result := twoSum(value.nums, value.target)

		for i := 0; i < len(result); i++ {
			if result[i] != value.solution[i] {
				t.Errorf("expected %v, got %v", value.solution, result)
			}
		}
	}
}
