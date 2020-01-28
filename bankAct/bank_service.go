package bankAct

type BankService interface {
	Withdraw(string, float64) error
	Deposit(string, float64) error
	Transfer(string, string, float64) error
	AccountExists(account string) bool

}
