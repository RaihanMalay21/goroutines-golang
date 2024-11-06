package belajar_golang_goroutines

import (
	"fmt"
	"time"
	"sync"
	"testing"
	"runtime"
)


func TestGomaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 *time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU: ", totalCpu)

	totalThred := runtime.GOMAXPROCS(-1)
	fmt.Println("Total thread: ", totalThred)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total goruotine: ", totalGoroutine)
}