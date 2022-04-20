package chapter4_condition

import (
	"fmt"
	"testing"
)

func TestIfCondition(t *testing.T) {
	if err := Connect(); err != nil {
		fmt.Println("test")
	} else {
		fmt.Println("ok")
	}
}

func Connect() error {
	return nil
}

func TestForCondition(t *testing.T) {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}
