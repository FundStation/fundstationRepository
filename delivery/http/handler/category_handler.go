package handler

import (
	"database/sql"
	"fmt"
	"github.com/FundStation2/bankAct/bank_repository"
	"github.com/FundStation2/bankAct/bank_service"
	"github.com/FundStation2/category"
	"github.com/FundStation2/form"
	"html/template"
	"net/http"
	"net/url"
)



type CategoryHandler struct {
	tmpl           *template.Template
	categoryService    category.CategoryService


}

func NewCategoryHandler(t *template.Template, cs category.CategoryService) *CategoryHandler {
	return &CategoryHandler{tmpl: t, categoryService: cs,}
}

func (cat *CategoryHandler) SelectMedicalCategory(w http.ResponseWriter, r *http.Request) {

	cato,err:=cat.categoryService.ViewCategory("medical")
	if err != nil{
		panic(err)
	}
	cat.tmpl.ExecuteTemplate(w,"listCategory.layout",cato)
}
func (cat *CategoryHandler) SelectWomenCategory(w http.ResponseWriter, r *http.Request) {

	cato,err:=cat.categoryService.ViewCategory("women")
	if err != nil{
		panic(err)
	}
	cat.tmpl.ExecuteTemplate(w,"listCategory.layout",cato)
}
func (cat *CategoryHandler) SelectEducationalCategory(w http.ResponseWriter, r *http.Request) {

	cato,err:=cat.categoryService.ViewCategory("education")
	if err != nil{
		panic(err)
	}
	cat.tmpl.ExecuteTemplate(w,"listCategory.layout",cato)
}
func (cat *CategoryHandler) SelectOrphanageCategory(w http.ResponseWriter, r *http.Request) {

	cato,err:=cat.categoryService.ViewCategory("orphanage")
	if err != nil{
		panic(err)
	}
	cat.tmpl.ExecuteTemplate(w,"listCategory.layout",cato)
}

func (cat *CategoryHandler) SelectReligiousCategory(w http.ResponseWriter, r *http.Request) {

	cato,err:=cat.categoryService.ViewCategory("religious")
	if err != nil{
		panic(err)
	}
	cat.tmpl.ExecuteTemplate(w,"listCategory.layout",cato)
}

func (cat *CategoryHandler) SelectOtherCategory(w http.ResponseWriter, r *http.Request) {

	cato,err:=cat.categoryService.ViewCategory("other")
	if err != nil{
		panic(err)
	}
	cat.tmpl.ExecuteTemplate(w,"listCategory.layout",cato)
}
var bankAct string
func (cat *CategoryHandler) SelectSpecificCategory(w http.ResponseWriter,r *http.Request){
	name:=r.FormValue("catName")
	cato,err:=cat.categoryService.ViewSpecificCategory(name)
	if err != nil{
		fmt.Println(err)
	}
	cat.tmpl.ExecuteTemplate(w,"specificEntity.layout",cato)
}

func (cat *CategoryHandler)Donation(w http.ResponseWriter,r *http.Request){

	cat.tmpl.ExecuteTemplate(w,"donation.layout",bankAcnt)
}

func (cat *CategoryHandler)DonateMoney(w http.ResponseWriter,r *http.Request) {


	dbconn, err := sql.Open("postgres", "postgres://postgres:password@localhost/funddb?sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer dbconn.Close()
	if r.Method == http.MethodGet {
		newDonForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
		}{
			Values:  nil,
			VErrors: nil,
		}
		cat.tmpl.ExecuteTemplate(w, "donation.layout", newDonForm)
	}
	bankRepo := bank_repository.NewPsqlBankRepository(dbconn)
	bankServ := bank_service.NewBankService(bankRepo)
	bch := NewBankAccountHandler(bankServ)

		err = r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		newDonForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		fmt.Println("Value", newDonForm.Values)
		newDonForm.Required("donorAccount", "amount")
		fmt.Println("newDonform", newDonForm)
		newDonForm.ExactLength("donorAccount", 13)
	//bch.transferMoney(w, r)
		if !newDonForm.Valid() {
			cat.tmpl.ExecuteTemplate(w, "donation.layout", newDonForm)
			fmt.Println("flifla")

			return
		}
		aExists := bch.bankAccountService.AccountExists(r.FormValue("donorAccount"))
		fmt.Println("aExixts", aExists)
		if aExists {
			newDonForm.VErrors.Add("donorAccount", "The account number you entered doesn't exist")
			cat.tmpl.ExecuteTemplate(w, "donation.layout", newDonForm)

			return
		}

		//if err != nil {
		//	cat.tmpl.ExecuteTemplate(w, "moneyTransfer.html", err)
		//}
		//tmpl.ExecuteTemplate(w, "moneyTransfer.html", nil)

	bch.transferMoney(w, r)
	//cat.tmpl.ExecuteTemplate(w,"donation.layout",bankAcnt)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
//}





