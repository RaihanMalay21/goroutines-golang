package belajar_golang_goroutines

import (
	"fmt"
	"time"
	"sync"
	"testing"
)

var cond = sync.NewCond(&sync.Mutex{})
var group = sync.WaitGroup{}

func WaitCondition(i int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", i)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {

	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// hanya mengijinkan satu dengan perintah signal
	// for j := 0; j < 10; j++ {
	// 	go func() {
	// 		time.Sleep(1 *time.Second)
	// 		cond.Signal()
	// 	}()
	// }

	// menjalankan semua perintah dengan broadcast
	go func() {
		time.Sleep(1 *time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}