package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct {
	balance int
	//mutex   *sync.Mutex
}

func (a *Account) Widthdraw(val int) {
	//a.mutex.Lock()
	a.balance -= val
	//a.mutex.Unlock()
}

func (a *Account) Deposit(val int) {
	//a.mutex.Lock()
	a.balance += val
	//a.mutex.Unlock()
}

func (a *Account) Balance() int {
	balance := a.balance
	return balance
}

var accounts []*Account
var globalLock *sync.Mutex

func Transfer(sender, receiver int, money int) {
	globalLock.Lock()
	accounts[sender].Widthdraw(money)
	accounts[receiver].Deposit(money)
	globalLock.Unlock()

	fmt.Println("Transfer", sender, receiver, money)
}

func GetTotalBalance() int {
	total := 0

	globalLock.Lock()
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	globalLock.Unlock()

	return total
}

func RandomTransfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer() {
	for {
		RandomTransfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n", GetTotalBalance())
}

func main() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 10000})
		//accounts = append(accounts, &Account{balance: 10000, mutex: &sync.Mutex{}})
	}

	globalLock = &sync.Mutex{}

	go func() {
		for {
			Transfer(0, 1, 100)
		}
	}()

	go func() {
		for {
			Transfer(1, 0, 100)
		}
	}()

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
