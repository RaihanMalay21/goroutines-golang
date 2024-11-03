package belajar_golang_goroutines


import (
	"fmt"
	"time"
	"sync"
	"testing"
)

func RunAsyncronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello World")
	time.Sleep(1 *time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncronous(group)
	}

	group.Wait()
	fmt.Println("selesai")
}