package belajar_golang_goroutines

import (
	"fmt"
	"time"
	"testing"
	"sync"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i <= 1000; i++ {
		go func() {
			for j := 0; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(7 *time.Second)
	fmt.Println("Counter = ", x)
}


type UserBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
} 

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance =  user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	user1.Change(-amount)

	user2.Lock()
	user2.Change(amount)

	user1.Unlock()
	user2.Unlock()
} 



func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name: "raihan",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name: "malay",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	fmt.Println("Total Balance ", user1.Name, " adalah ", user1.Balance)
	fmt.Println("Total Balance ", user2.Name, " adalah ", user2.Balance)
}