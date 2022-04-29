package chapter12_benchmark

import (
	"sync"
	"sync/atomic"
	"testing"
)

var n int32

func addSyncByAtomic(delta int32) int32 {
	return atomic.AddInt32(&n, delta)
}

func getSyncByAtomic() int32 {
	return atomic.LoadInt32(&n)
}

var n2 int
var mu sync.Mutex

func addSyncByMutex(delta int) {
	mu.Lock()
	n2 += delta
	mu.Unlock()
}

func readSyncByMutex() int {
	var n int
	mu.Lock()
	n = n2
	mu.Unlock()

	return n
}

func BenchmarkAddSyncByAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addSyncByAtomic(1)
		}
	})
}

func BenchmarkReadSyncByAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			getSyncByAtomic()
		}
	})
}
func BenchmarkAddSyncByMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addSyncByMutex(1)
		}
	})
}

func BenchmarkReadSyncByMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			readSyncByMutex()
		}
	})
}
