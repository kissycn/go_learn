package array

import (
	"fmt"
	"testing"
)

func TestRotateArray(t *testing.T) {
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	nums := []int{-1}
	rotate(nums, 1)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	//k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
