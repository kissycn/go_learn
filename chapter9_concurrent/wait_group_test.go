package chapter9_concurrent

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	urls := []string{
		"http://www.qq.com",
		"http://www.baidu.com",
		"http://www.taobao.com",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			fmt.Println(u)
		}(url)
	}

	wg.Wait()
	fmt.Println("all end")
}
