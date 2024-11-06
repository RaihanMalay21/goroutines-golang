package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"sync/atomic"
)

func TestAtomic(t *testing.T) {
	var group = sync.WaitGroup{}
	var x int64 = 0

	for i := 1; i <= 1000; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
		}()
	}

	group.Wait()

	fmt.Println("Counter = ", x)
}