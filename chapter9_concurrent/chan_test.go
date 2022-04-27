package chapter9_concurrent

import (
	"fmt"
	"testing"
)

func TestChanInit(t *testing.T) {
	ch1 := make(chan int)
	go func() {
		fmt.Println("start goroutine")
		ch1 <- 0
		fmt.Println("end goroutine")
	}()

	<-ch1
	fmt.Println("all done")
}

func printer(ch chan int) {
	for {
		data := <-ch
		if data == 0 {
			break
		}

		fmt.Println(data)
	}
	ch <- 0
}

func TestPrinter(t *testing.T) {
	ch1 := make(chan int)

	go printer(ch1)
	for i := 1; i < 10; i++ {
		ch1 <- i
	}
	ch1 <- 0
	<-ch1
}

func TestBufferChan(t *testing.T) {
	ch1 := make(chan int, 3)
	fmt.Println(len(ch1))
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	fmt.Println(len(ch1))
}
