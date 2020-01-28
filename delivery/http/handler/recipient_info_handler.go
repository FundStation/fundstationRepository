package handler

import (
	"database/sql"
	"fmt"
	"github.com/FundStation/bankAct/bank_repository"
	"github.com/FundStation/bankAct/bank_service"
	"github.com/FundStation/form"
	"github.com/FundStation/models"
	"github.com/FundStation/recipientInfo"
	"github.com/FundStation/tokens"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"html/template"
)

type RecipientInfoHandler struct {
	tmpl           *template.Template
	recpInfoService    recipientInfo.RecipientInfoService
	csrfSignKey    []byte

}

func NewRecipientInfoHandler(t *template.Template, rs recipientInfo.RecipientInfoService,csKey []byte) *RecipientInfoHandler {
	return &RecipientInfoHandler{tmpl: t, recpInfoService: rs,csrfSignKey:csKey}
}
func (rch *RecipientInfoHandler) RecipientInfo(w http.ResponseWriter, r *http.Request) {
	token, err := tokens.CSRFToken(rch.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	newCatForm := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	rch.tmpl.ExecuteTemplate(w,"recipientInfo1.layout",newCatForm)

}
func (rch *RecipientInfoHandler) CreateRecipientInfo(w http.ResponseWriter, r *http.Request) {
	token, err := tokens.CSRFToken(rch.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		newCatForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		rch.tmpl.ExecuteTemplate(w, "recipientInfo1.layout", newCatForm)
	}

	if r.Method == http.MethodPost {

		err := r.ParseMultipartForm(1024)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		fmt.Println("accountNUm",r.FormValue("accountNo"))
		newDonForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		//newDonForm.Required("accountNo", "description")
		//newDonForm.ExactLength("accountNo", 13)

		newDonForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !newDonForm.Valid() {
			rch.tmpl.ExecuteTemplate(w, "recipientInfo1.layout", newDonForm)
			fmt.Println("why")

			return
		}
		dbconn, err := sql.Open("postgres", "postgres://postgres:password@localhost/funddb?sslmode=disable")

		if err != nil {
			panic(err)
		}
		defer dbconn.Close()
		bankRepo := bank_repository.NewPsqlBankRepository(dbconn)
		bankServ := bank_service.NewBankService(bankRepo)
		bch := NewBankAccountHandler(bankServ)
		aaExists := bch.bankAccountService.AccountExists(r.FormValue("accountNo"))
		if aaExists {
			newDonForm.VErrors.Add("accountNo", "The account number you entered doesn't exist")
			rch.tmpl.ExecuteTemplate(w, "recipientInfo1.layout", newDonForm)
			return
		}
		aExists := rch.recpInfoService.AccountExistsInfo(r.FormValue("accountNo"))
		if aExists {
			newDonForm.VErrors.Add("accountNo", "Account number already exists in the previously submitted forms")
			rch.tmpl.ExecuteTemplate(w, "recipientInfo1.layout", newDonForm)
			return
		}


		mf, fh, _ := r.FormFile("recpimg")

			defer mf.Close()

		mff, fhh, _ := r.FormFile("attachment")
		defer mff.Close()

			writeImageFile(&mf, fh.Filename)
		writeFile(&mff,fhh.Filename)
		goal, _ := strconv.ParseFloat(r.FormValue("goal"), 64)
			recipientInfo := models.RecipientInfo{
				Image:       fh.Filename,
				Description: r.FormValue("description"),
				Attachment:  fhh.Filename,
				Recipient:   &recp,
				Date:        time.Now(),
				BankAccount: models.BankAccount{
					AccountNo: r.FormValue("accountNo"),
				},
				Goal:         goal,
			}

			fmt.Println("here",recp)

			erro := rch.recpInfoService.CreateRecipientInfo(recipientInfo)

			if erro != nil {
				fmt.Println(erro)
				panic(err)

			}
			//http.Redirect(w, r, "/home", http.StatusSeeOther)
			rch.tmpl.ExecuteTemplate(w,"recipientPage.layout",recipientInfo)

		}

	}
func(rch *RecipientInfoHandler) ShowApproved(w http.ResponseWriter, r *http.Request){
	rinfos,err:=rch.recpInfoService.SelectApproved()
	fmt.Println("Rinfos:",rinfos)
	if err != nil{
		rch.tmpl.ExecuteTemplate(w,"category.layout",nil)
		return
	}

	rch.tmpl.ExecuteTemplate(w,"category.layout",rinfos[0:2])
}

func(rch *RecipientInfoHandler) SeeAll(w http.ResponseWriter, r *http.Request){
	rinfos,err:=rch.recpInfoService.SelectApproved()
	if err != nil{
		rch.tmpl.ExecuteTemplate(w,"allindividual.layout",nil)
		return
	}
	rch.tmpl.ExecuteTemplate(w,"allindividual.layout",rinfos)
}
var Path string
func(rch *RecipientInfoHandler) SeeIndividual(w http.ResponseWriter, r *http.Request){
	Path=r.URL.Path
	fmt.Println(Path)
	idRaw:=strings.Trim(Path,"/")
	id,_:=strconv.Atoi(idRaw)

	rinfo,err:=rch.recpInfoService.SelectIndividualById(id)
	fmt.Println(rinfo)
	if err != nil{
		rch.tmpl.ExecuteTemplate(w,"individual.layout",nil)
		return
	}
	rch.tmpl.ExecuteTemplate(w,"individual.layout",rinfo)
}
func (rch *RecipientInfoHandler) ViewSpecificRecipientInfo(w http.ResponseWriter, r *http.Request) {
	rNo, _ := strconv.Atoi(r.FormValue("rNo"))

	recp, err := rch.recpInfoService.ViewSpecificRecipientInfo(rNo)

	if err != nil {
		panic(err)
	}

	rch.tmpl.ExecuteTemplate(w, "specificRecipient.html", recp)
}


func(rch *RecipientInfoHandler) ApproveRecipient(w http.ResponseWriter, r *http.Request){
	riNo, _ := strconv.Atoi(r.FormValue("recipientInfoId"))
	err:=rch.recpInfoService.ApproveRecipientInfo(riNo)

	if err != nil{
		panic(err)
	}

	http.Redirect(w,r,"/admin",http.StatusSeeOther)
}


	func writeImageFile(mf *multipart.File, fname string) {

		wd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		path := filepath.Join(wd, "web", "assets", "img", fname)
		image, err := os.Create(path)

		if err != nil {
			panic(err)
		}
		defer image.Close()
		io.Copy(image, *mf)
	}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "web", "assets", "attachment", fname)
	attachment, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer attachment.Close()
	io.Copy(attachment, *mf)
}

