// Posting and transtation logic
package core

import "time"

type Posting struct {
	AccountId string
	Amount    float64 // Positive for Debit, Negative for Credit? or explict db/cr fields
}

type Transaction struct {
	Date        time.Time
	Description string
	Postings    []Posting // must sum to 0 to balance
}
