package account

import "sync"

const testVersion = 1

// SafeCounter is safe to use concurrently.

type SafeCounter struct {
	v   int64
	mux sync.Mutex
}

// Increase value and returns the current value of the counter
func (c *SafeCounter) Value() int64 {
	c.mux.Lock()
	c.v++
	defer c.mux.Unlock()
	return c.v
}

var gid SafeCounter = SafeCounter{v: 0}

type Account struct {
	id      int64
	balance int64
	closed  bool
	mux     sync.RWMutex
}

func Open(initalDeposit int64) *Account {
	if initalDeposit < 0 {
		return nil
	}
	var accid int64 = gid.Value()
	a := Account{id: accid, balance: initalDeposit, closed: false}
	return &a
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mux.Lock()

	if a.closed {
		payout = 0
		ok = false
	} else {
		a.closed = true
		payout = a.balance
		ok = true
	}
	defer a.mux.Unlock()
	return payout, ok

}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mux.RLock()

	if a.closed {
		balance = 0
		ok = false
	} else {
		balance = a.balance
		ok = true
	}
	defer a.mux.RUnlock()
	return balance, ok

}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mux.Lock()

	if a.closed {
		newBalance = 0
		ok = false
	} else {
		if a.balance+amount < 0 {
			ok = false
		} else {
			a.balance = a.balance + amount
			ok = true
		}
		newBalance = a.balance
	}
	defer a.mux.Unlock()
	return newBalance, ok

}
