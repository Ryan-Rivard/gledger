package main

import "fmt"

// Transaction represents a classic double-entry record.
type Transaction struct {
	Date        string
	Description string
	DebitAcc    string
	CreditAcc   string
	AmountCents int64
}

// Ledger holds an index array of all valid transactions.
type Ledger struct {
	Transactions []Transaction
}

// ValidateEntry verifies basic constraints before allowing a transaction to be stored.
func (l *Ledger) ValidateEntry(date, desc, debit, credit string, amountCents int64) (*Transaction, error) {
	if amountCents <= 0 {
		return nil, fmt.Errorf("amount must be a positive number of cents")
	}
	if len(debit) == 0 || len(credit) == 0 {
		return nil, fmt.Errorf("both debit and credit accounts must be specified")
	}
	if len(date) == 0 || len(desc) == 0 {
		return nil, fmt.Errorf("date and description cannot be empty")
	}

	return &Transaction{
		Date:        date,
		Description: desc,
		DebitAcc:    debit,
		CreditAcc:   credit,
		AmountCents: amountCents,
	}, nil
}
