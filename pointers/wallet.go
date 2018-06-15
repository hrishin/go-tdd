package main

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposite(amount Bitcoin) {
	w.balance = w.balance + amount
}

func (w *Wallet) Balance() Bitcoin {
	return (w.balance)
}
