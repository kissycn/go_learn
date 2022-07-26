package array

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSums(t *testing.T) {
	nums := []int{-2, 0, 0, 2, 2}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for k := 0; k < len(nums)-2; k++ {
		if nums[k] > 0 {
			//若nums[k]大于0，则后面的数字也是大于零（排序后是递增的）
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue //nums[k]值重复了，去重
		}
		left, right := k+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right] + nums[k]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				res = append(res, []int{nums[left], nums[right], nums[k]})
				for right > left && nums[right] == nums[right-1] {
					right--
				}
				for right > left && nums[left] == nums[left+1] {
					left++
				}
				left++
				right--
			}
		}
	}
	return res
}

// 两数之和，采用hash求解
func threeSum2(nums []int) [][]int {
	res := [][]int{}
	m := map[int][]int{}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			c := 0 - nums[i] - nums[j]
			if val, ok := m[c]; ok {
				res = append(res, val)
			} else {
				m[c] = []int{nums[i], nums[j], c}
			}
			// a + b + c = 0 ----> a + b = -c
			//if val, ok := m[-(nums[i] + nums[j])]; ok {
			//	res = append(res, val)
			//} else {
			//	// a + b + c = 0 ----> c = 0 - a - b
			//	c := 0 - nums[i] - nums[j]
			//	m[c] = []int{nums[i], nums[j], c}
			//}
		}
	}
	return res
}

// 暴力求解法，三层遍历实践复杂度为：O(n^3)
func threeSum1(nums []int) [][]int {
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
