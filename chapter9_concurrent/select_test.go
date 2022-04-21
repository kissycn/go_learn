package chapter9_concurrent

import (
	"fmt"
	"testing"
	"time"
)

func channel1(ch1 chan int) {
	for i := 0; ; i++ {
		ch1 <- i + 1
	}
}

func channel2(ch2 chan int) {
	for i := 0; ; i++ {
		ch2 <- i + 2
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v1 := <-ch1:
			fmt.Println("channel-1:", v1)
		case v2 := <-ch2:
			fmt.Println("channel-2", v2)
		}
	}
}

func TestSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go channel1(ch1)
	go channel2(ch2)
	go suck(ch1, ch2)
	time.Sleep(time.Second)
}
