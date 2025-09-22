package main

import "time"

type Account struct {
	AccountID      int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
	CreatedAt      string `json:"created_at"`
}

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

type Transaction struct {
	TransactionID   int64   `json:"transaction_id"`
	AccountID       int64   `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
	EventDate       string  `json:"event_date"`
}

type CreateTransactionRequest struct {
	AccountID       int64   `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

const (
	OpCashPurchase        = 1
	OpInstallmentPurchase = 2
	OpWithdrawal          = 3
	OpPayment             = 4
)

func NowISO() string {
	return time.Now().Format(time.RFC3339Nano)
}
