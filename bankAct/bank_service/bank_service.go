package bank_service

import (
	"errors"
	"github.com/FundStation/bankAct"
)

type BankService struct {
	bankRepo bankAct.BankRepository
}

func NewBankService(BankRepo bankAct.BankRepository) *BankService {
	return &BankService{bankRepo: BankRepo}
}
func (bs BankService) Withdraw(accountNo string, transAmount float64) error {
	account, err := bs.bankRepo.SelectAccountNo(accountNo)
	if err != nil {
		return err

	}
	if transAmount > account.CurrentBalance {
		return errors.New("Insufficent Amount")
	}
	account.CurrentBalance -= transAmount
	err = bs.bankRepo.UpdateBalance(accountNo, account.CurrentBalance)
	if err != nil {
		return err
	}
	return nil

}

func (bs BankService) Deposit(accountNo string, transAmount float64) error {

	account, err := bs.bankRepo.SelectAccountNo(accountNo)
	if err != nil {
		return err

	}
	account.CurrentBalance += transAmount
	err = bs.bankRepo.UpdateBalance(accountNo, account.CurrentBalance)
	if err != nil {
		return errors.New("Could not transfer the given amount!")
	}
	return nil

}

func (bs BankService) Transfer(donAccountNo string, recAccountNo string, transAmount float64) error {
	err := bs.Withdraw(donAccountNo, transAmount)
	if err != nil {
		//panic(err)
		return err
	}
	err = bs.Deposit(recAccountNo, transAmount)
	if err != nil {
		//panic(err)
		return err
	}

	return nil

}
