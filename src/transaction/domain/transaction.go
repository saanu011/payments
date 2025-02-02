package domain

import "time"

type Transaction struct {
	ID              string        `db:"id"`
	AccountID       string        `db:"account_id"`
	OperationTypeID OperationType `db:"operation_type_id"`
	Amount          float64       `db:"amount"`
	CreatedAt       time.Time     `db:"created_at"`
	UpdatedAt       time.Time     `db:"updated_at"`
}

type OperationType int

const (
	Normal        OperationType = 1
	Installment   OperationType = 2
	Withdrawal    OperationType = 3
	CreditVoucher OperationType = 4
)
