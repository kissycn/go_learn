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
