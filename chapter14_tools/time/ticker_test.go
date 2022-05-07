package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTickerTest(t *testing.T) {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

func TestTimerAndTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(time.Second * 10)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}
