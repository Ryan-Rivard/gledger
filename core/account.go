// Account types and chart definitions
package core

type AccountType string

const (
	Asset     AccountType = "Asset"
	Liability AccountType = "Liability"
	Equity    AccountType = "Equity"
)

type Account struct {
	Id   string
	Name string
	Type AccountType
}
