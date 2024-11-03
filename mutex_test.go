package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"time"
	"testing"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(3 *time.Second)
	fmt.Println("Counter = ", x)
}

type BackAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BackAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BackAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BackAccount{}

	for i := 0; i <= 100; i++ {
		go func() {
			for i := 0; i <= 100; i++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 *time.Second)
	fmt.Println("Total Balance ", account.Balance)
}