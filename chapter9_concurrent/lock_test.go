package chapter9_concurrent

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	count int32
	wg    sync.WaitGroup
)

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		atomic.AddInt32(&count, 1)
		runtime.Gosched()
	}
}

func TestCount(t *testing.T) {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}

var (
	counter1 int
	lock     sync.Mutex
)

func SetData(c int) {
	defer lock.Unlock()

	lock.Lock()
	counter1 = c
}

func GetData() int {
	defer lock.Unlock()

	lock.Lock()
	return counter1

}
