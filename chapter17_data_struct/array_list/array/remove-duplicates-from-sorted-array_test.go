package array

import (
	"fmt"
	"testing"
)

func TestDuplicatesArray(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates(nums)
	fmt.Println(n)
}

// 双指针法，left指针放不重复元素的位置，right进行遍历
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 1
	}
	left, right := 1, 1

	for right < n {
		// 不相等说明遇到了不重复的元素，把该元素挪到不重复位置上去，并更新指针
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	//nums1 := nums[:left]
	//fmt.Println(nums1)
	return left
}

func removeDuplicates1(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	left, right, n := 0, 1, len(nums)
	for right < n {
		if nums[left] == nums[right] {
			nums = append(nums[:right], nums[right+1:]...)
			n--
		} else {
			left++
			right++
		}
	}
	return n
}
