package array

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4, 1}
	target := 6
	fmt.Println(twoSum1(nums, target))
}

func twoSum(nums []int, target int) []int {
	a := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				a = append(a, i)
				a = append(a, j)
				return a
			}
		}
	}
	return []int{}
}

func twoSum1(nums []int, target int) []int {
	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			if v+nums[k] == target {
				return []int{i, k}
			}
		}
	}
	return []int{}
}
