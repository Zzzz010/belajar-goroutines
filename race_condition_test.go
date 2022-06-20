package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
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

	time.Sleep(3 * time.Second)

	fmt.Println("Counter = ", x)

}

// mengatasi measalah Race Condition

func TestSyncMutex(t *testing.T) {
	x := 0
	var Mutex sync.Mutex
	for i := 0; i <= 1000; i++ {
		go func() {
			for j := 0; j <= 100; j++ {
				Mutex.Lock()
				x = x + 1
				Mutex.Unlock()
			}
		}()
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Counter = ", x)

}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.Lock()
	balance := account.Balance
	account.RWMutex.Unlock()
	return balance
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance : ", account.GetBalance())
}

// --------------------------------------------

// Proble Deadlock

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Trasfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Bima",
		Balance: 500000,
	}
	user2 := UserBalance{
		Name:    "Arjuna",
		Balance: 100,
	}

	go Trasfer(&user1, &user2, 100000)
	go Trasfer(&user2, &user1, 200000)

	time.Sleep(10 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)

	// go Trasfer(&user1, &user2, 1000)
	// go Trasfer(&user1, &user2, 1000)

}
