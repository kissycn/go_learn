package array

import (
	"fmt"
	"testing"
)

// 采用数组数组嵌套遍历，然后去除最大值进行
// 时间复杂度为：O(n^2)
func TestArrayIter(t *testing.T) {
	height := [9]int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	max := 0
	// 双重循环遍历数组的左右边界且不重复
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area := min(height[i], height[j]) * (j - i)
			if area > max {
				max = area
			}
		}
	}

	fmt.Println(max)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 双指针，两边向中间收敛，收敛策略为：每次总是移动较小的那个
func TestArrayByIndex(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	left, right := 0, len(height)-1
	area, max := 0, 0

	for left < right {
		area = min(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}

		// 如果是 <= 则会重复算一次
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	fmt.Println(max)
}
