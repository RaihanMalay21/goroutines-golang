package belajar_golang_goroutines

import (
	"fmt"
	"time"
	"testing"
)

func HelloWord() {
	fmt.Println("Hello World")
}

func TestCreateGoroutines(t *testing.T) {
	go HelloWord()
	fmt.Println("bayar")

	time.Sleep(1 * time.Second)
}


func displayNumber(i int) {
	fmt.Println("display", i)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go displayNumber(i)
	}

	time.Sleep(1 * time.Second)
}