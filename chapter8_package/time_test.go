package chapter8_package

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()

	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Minute())

	fmt.Println(now.Unix())

	fmt.Println(now)
}
