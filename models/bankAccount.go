package models

type BankAccount struct {
	AccountNo      string  `json:"accountno"`
	CurrentBalance float64
}
