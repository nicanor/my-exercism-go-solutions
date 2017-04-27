package account

import "sync"

const testVersion = 1

type Account struct {
	open    bool
	balance int64
	sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, open: true}
}

func (a *Account) Balance() (balance int64, ok bool) {
	return a.balance, a.open
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	defer a.Unlock()
	a.Lock()
	var total int64 = a.balance + amount
	if !a.open || total < 0 {
		return 0, false
	}
	a.balance = total
	return a.balance, true
}

func (a *Account) Close() (payout int64, ok bool) {
	defer a.Unlock()
	a.Lock()
	if !a.open {
		return 0, false
	}
	a.open = false
	return a.balance, true
}
