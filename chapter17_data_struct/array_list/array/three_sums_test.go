package array

import (
	"fmt"
	"testing"
)

func TestThreeSums(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

// 暴力求解法，三层遍历实践复杂度为：O(n^3)
func threeSum(nums []int) [][]int {
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return res
}
