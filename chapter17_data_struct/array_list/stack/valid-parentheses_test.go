package stack

import (
	"fmt"
	"testing"
)

func TestParentheses(t *testing.T) {
	fmt.Println(isValid("(([{}]))"))
}

// 每次遇到左括号，把右括号推入栈，这样直接比较括号不用查询
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	stack := []int32{}
	for _, b := range s {
		switch b {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		default:
			if len(stack) == 0 || stack[len(stack)-1] != b {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
