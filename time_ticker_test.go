package belajar_golang_goroutines

import (
	"fmt"
	"time"
	"testing"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case timer := <- ticker.C:
			fmt.Println(timer)
		case <- done: 
			return
		}
	}

}

func TestTick(t *testing.T) {
	
}