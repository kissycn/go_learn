package array

import (
	"testing"
)

func TestArray(t *testing.T) {
	a := []int{0}
	moveZeroes(a)
}

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)

	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func moveZeroes1(nums []int) {
	if nums == nil {
		return
	}

	newArr := make([]int, len(nums))
	front := 0
	back := len(nums) - 1
	for _, v := range nums {
		if v == 0 {
			newArr[back] = 0
			back--
		} else {
			newArr[front] = v
			front++
		}
	}
}
