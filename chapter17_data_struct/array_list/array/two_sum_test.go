package array

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4, 1}
	target := 6
	fmt.Println(twoSumByMap(nums, target))
}

func twoSumByMap(nums []int, target int) []int {
	// 为了减少二层循环遍历，将第二次循环的val作为map的key存下来，空间换时间
	m := map[int]int{}
	for i, v := range nums {
		if val, ok := m[target-v]; ok {
			return []int{val, i}
		}
		// 一定要注意key是对应数组的val
		m[v] = i
	}

	return []int{}
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
