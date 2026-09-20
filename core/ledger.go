// Core ledger state and transaction matching
package core

type Ledger struct {
	Accounts     map[string]*Account
	Transactions []Transaction
}
