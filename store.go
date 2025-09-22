package main

import (
	"errors"
	"sync"
)

var (
	ErrAccountNotFound   = errors.New("account not found")
	ErrInvalidOperation  = errors.New("invalid operation type")
	ErrInvalidAmountSign = errors.New("amount sign invalid for operation type")
)

type Store struct {
	mu           sync.Mutex
	nextAccount  int64
	nextTrans    int64
	accounts     map[int64]Account
	transactions map[int64]Transaction
}

func NewStore() *Store {
	return &Store{
		nextAccount:  1,
		nextTrans:    1,
		accounts:     make(map[int64]Account),
		transactions: make(map[int64]Transaction),
	}
}

func (s *Store) CreateAccount(doc string) Account {
	s.mu.Lock()
	defer s.mu.Unlock()

	a := Account{
		AccountID:      s.nextAccount,
		DocumentNumber: doc,
		CreatedAt:      NowISO(),
	}
	s.accounts[s.nextAccount] = a
	s.nextAccount++
	return a
}

func (s *Store) GetAcount(id int64) (Account, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	a, ok := s.accounts[id]
	if !ok {
		return Account{}, ErrAccountNotFound
	}
	return a, nil
}

// validate amoutn sign based on operation type
// purchases(1,2) and withdrawal(3) must have negative amount,
// payment(4) must have positive amount
func validateAmountSign(opType int, amount float64) error {
	switch opType {
	case OpCashPurchase, OpInstallmentPurchase, OpWithdrawal:
		if amount >= 0 {
			return ErrInvalidAmountSign
		}
		return nil
	case OpPayment:
		if amount <= 0 {
			return ErrInvalidAmountSign
		}
		return nil
	default:
		return ErrInvalidOperation
	}
}

func (s *Store) CreateTransaction(req CreateTransactionRequest) (Transaction, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	//account must exist
	// if _, ok := s.accounts(req.AccountID); !ok {
	// 	return Transaction{}, ErrAccountNotFound
	// }

	//validate op type & amount sign
	if err := validateAmountSign(req.OperationTypeID, req.Amount); err != nil {
		return Transaction{}, err
	}

	t := Transaction{
		TransactionID:   s.nextTrans,
		AccountID:       req.AccountID,
		OperationTypeID: req.OperationTypeID,
		Amount:          req.Amount,
		EventDate:       NowISO(),
	}
	s.transactions[s.nextTrans] = t
	s.nextTrans++
	return t, nil
}
