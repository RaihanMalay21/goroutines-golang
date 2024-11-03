package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	// "time"
	"testing"
)

func TestPool(t *testing.T) {
	var group = sync.WaitGroup{}
	pool := sync.Pool{
		// default values
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Raihan")
	pool.Put("Malay")
	pool.Put("Ganteng And Kaya")


	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			data := pool.Get()
			fmt.Println(data)
		}()
	}

	group.Wait()
	fmt.Println("selesai")
} 