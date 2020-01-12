package handler

import (
	"fmt"
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

		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		newDonForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		//newDonForm.Required("accountNo", "description")
		//newDonForm.MinLength("accountNo", 13)

		newDonForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !newDonForm.Valid() {
			rch.tmpl.ExecuteTemplate(w, "recipientInfo1.layout", newDonForm)
			fmt.Println("why")

			return
		}
		aExists := rch.recpInfoService.AccountExistsInfo(r.FormValue("accountNo"))
		if aExists {
			newDonForm.VErrors.Add("accountNo", "Account Number Already Exists")
			rch.tmpl.ExecuteTemplate(w, "recipientInfo1.layout", newDonForm)
			return
		}


			mf, fh, err := r.FormFile("recpimg")
			if err != nil {
				panic(err)
			}
			defer mf.Close()

		mff, fhh, err := r.FormFile("attachment")
		if err != nil {
			panic(err)
		}
		defer mff.Close()

			writeImageFile(&mf, fh.Filename)
			recipientInfo := models.RecipientInfo{
				Image:       fh.Filename,
				Description: r.FormValue("description"),
				Attachment:  fhh.Filename,
				Recipient:   &recp,
				Date:        time.Now(),
				BankAccount: models.BankAccount{
					AccountNo: r.FormValue("accountNo"),
				},
			}

			fmt.Println(recp)

			erro := rch.recpInfoService.CreateRecipientInfo(recipientInfo)

			if erro != nil {
				panic(err)

			}
			http.Redirect(w, r, "/home", http.StatusSeeOther)

		}

	}
func(rch *RecipientInfoHandler) SelectApproved(w http.ResponseWriter, r *http.Request){
	rinfos,err:=rch.recpInfoService.SelectApproved()
	if err != nil{
		tmpl.ExecuteTemplate(w,"category.layout",nil)
		return
	}
	tmpl.ExecuteTemplate(w,"category.layout",rinfos)
}
func (rch *RecipientInfoHandler) ViewSpecificRecipientInfo(w http.ResponseWriter, r *http.Request) {
	rNo, _ := strconv.Atoi(r.FormValue("rNo"))

	recp, err := rch.recpInfoService.ViewSpecificRecipientInfo(rNo)

	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "specificRecipient.html", recp)
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