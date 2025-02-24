package bank

import (
	"fmt"
)

// 1
type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

func (s *savingsAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid deposit amount")
	}
	s.balance += amount
	return nil
}
func (s *savingsAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid withdraw amount")
	}
	if amount > s.balance {
		return &InsufficientFundsError{Requested: amount, Balance: s.balance}
	}
	s.balance -= amount
	return nil
}

func (s *savingsAccount) GetBalance() float64 {
	return s.balance
}

// 2
type savingsAccount struct {
	balance float64
}

func NewSavingsAccount(initialBalance float64) *savingsAccount {
	return &savingsAccount{balance: initialBalance}
}

func NewInterestAccount(initialBalance, interestRate float64) *interestAccount {
	return &interestAccount{
		savingsAccount: savingsAccount{balance: initialBalance},
		interestRate:   interestRate,
	}
}

// 3
type interestAccount struct {
	savingsAccount
	interestRate float64
}

func (ia *interestAccount) AddInterest() {
	interest := ia.balance * ia.interestRate
	ia.balance += interest
}

// 4
func Transfer(from, to Account, amount float64) error {
	err := from.Withdraw(amount)
	if err != nil {
		return err
	}

	err = to.Deposit(amount)
	if err != nil {
		_ = from.Deposit(amount)
		return err
	}

	return nil
}

func PrintBalance(a Account) {
	fmt.Println(a.GetBalance())
}

// 5

type InsufficientFundsError struct {
	Requested float64
	Balance   float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("ยอดเงินไม่เพียงพอในการถอน: ต้องการ %.2f แต่มีแค่ %.2f", e.Requested, e.Balance)
}

// 6
func RiskyBankOperation(n float64) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("เกิด Panic แล้วทำการ Recover สำเร็จ: %v\n", r)
		}
	}()
	if n < 0 {
		panic("เกิด Panic แล้วทำการ Recover สำเร็จ:")
	}
	fmt.Println("บรรทัดนี้จะไม่ถูกพิมพ์")
}

// 7

//8
