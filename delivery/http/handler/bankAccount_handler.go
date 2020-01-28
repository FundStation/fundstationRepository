package handler

import (
	"fmt"
	bankAct2 "github.com/FundStation2/bankAct"
	"strconv"

	"net/http"
)

type BankAccountHandler struct {
	bankAccountService  bankAct2.BankService


}
func NewBankAccountHandler( bs bankAct2.BankService) *BankAccountHandler {
	return &BankAccountHandler{ bankAccountService: bs}
}
var transMoney float64
var donAct string
var recAct string
func (bch *BankAccountHandler) transferMoney(w http.ResponseWriter, r *http.Request) {
	//tmpl := template.Must(template.ParseGlob("web/template/*"))
	if r.Method == http.MethodPost {
		transMoney, _ = strconv.ParseFloat(r.FormValue("amount"), 64)
		donAct = r.FormValue("donorAccount")

		recAct = r.FormValue("recpAccount")
		err :=bch.bankAccountService.Transfer(donAct, recAct, transMoney)
		fmt.Println("form value",r.FormValue("donorAccount"))
		if err != nil {
			//tmpl.ExecuteTemplate(w, "moneyTransfer.html", err)
		}
		//tmpl.ExecuteTemplate(w, "moneyTransfer.html", nil)
	}else {

		//tmpl.ExecuteTemplate(w, "moneyTransfer.html", nil)
	}

}
