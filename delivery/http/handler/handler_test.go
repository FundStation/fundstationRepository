package handler

import (
	"bytes"
	"fmt"
	"github.com/FundStation2/category/category_repository"
	"github.com/FundStation2/category/category_service"
	"github.com/FundStation2/donor/donor_repository"
	"github.com/FundStation2/donor/donor_service"
	"github.com/FundStation2/models"
	"github.com/FundStation2/recipient/recipient_repository"
	"github.com/FundStation2/recipient/recipient_service"
	"github.com/FundStation2/recipientInfo/recipientInfo_repository"
	"github.com/FundStation2/recipientInfo/recipientInfo_service"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHome(t *testing.T){
	w := httptest.NewRecorder()
	r,err := http.NewRequest("GET","/home",nil)
	if err != nil{
		t.Fatal(err)
	}
	Home(w,r)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK{
		t.Errorf("want %d, got %d",http.StatusOK,resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		t.Fatal(err)
	}
	if testing.Short(){
		t.Skip("Skipping long-running test in short mode")
	}
	if string(body) != "Home"{
		t.Errorf("want the body to contain the word %q","Home")
	}

}
func TestDonorSignup(t *testing.T) {
	tmpl := template.Must(template.ParseGlob("../../../web/template/*"))
	donorRepo := donor_repository.NewMockDonorRepo(nil)
	donorServ := donor_service.NewDonorService(donorRepo)

	donorHandler := NewDonorHandler (tmpl, donorServ,nil,nil,nil,nil)



	mux := http.NewServeMux()
	mux.HandleFunc("/donor/signup", donorHandler.DonorSignup)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("firstName", models.DonorMock.FirstName)
	form.Add("lastName", models.DonorMock.LastName)
	form.Add("add", models.DonorMock.Address)
	form.Add("userName", models.DonorMock.Username)
	form.Add("password", models.DonorMock.Password)
	form.Add("phone", models.DonorMock.PhoneNumber)
	form.Add("email", models.DonorMock.EmailAddress)



	resp, err := tc.PostForm(sURL+"/donor/signup", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("mockyuser")) {
		t.Errorf("want body to contain %q", body)
	}

}
func TestRecipientSignup(t *testing.T) {
	tmpl := template.Must(template.ParseGlob("../../../web/template/*"))
	recipientRepo := recipient_repository.NewMockRecipientRepo(nil)
	recipientServ := recipient_service.NewRecipientService(recipientRepo)

	recipientHandler := NewRecipientHandler (tmpl, recipientServ,nil,nil,nil,nil)



	mux := http.NewServeMux()
	mux.HandleFunc("/recipient/signup", recipientHandler.RecipientSignup)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("firstName", models.RecipientMock.FirstName)
	form.Add("lastName", models.RecipientMock.LastName)
	form.Add("add", models.RecipientMock.Address)
	form.Add("userName", models.RecipientMock.Username)
	form.Add("password", models.RecipientMock.Password)
	form.Add("phone", models.RecipientMock.PhoneNumber)
	form.Add("email", models.RecipientMock.EmailAddress)



	resp, err := tc.PostForm(sURL+"/recipient/signup", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("mockyuser")) {
		t.Errorf("want body to contain %q", body)
	}

}
func TestCreateRecipientInfo(t *testing.T){
	tmpl := template.Must(template.ParseGlob("../../../web/template/*"))
	recipientInfoRepo := recipientInfo_repository.NewMockRecipientInfoRepository(nil)
	recipientInfoServ := recipientInfo_service.NewRecipientInfoService(recipientInfoRepo)

	recipientInfoHandler := NewRecipientInfoHandler (tmpl, recipientInfoServ,nil,)



	mux := http.NewServeMux()
	mux.HandleFunc("/recipientInfo/create", recipientInfoHandler.CreateRecipientInfo)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL
	//ID          int               `json:"id"`
	//Image       string			  `json:"image"`
	//Description string			  `json:"description"`
	//Attachment  string			  `json:"attachment"`
	//Recipient   *Recipient		  `json:"recipient_id"`
	//Date        time.Time		  `json:"submitteddate"`
	//BankAccount BankAccount       `json:"account"`
	//Goal        float64			  `json:"goal"`
	form := url.Values{}
	//mage,description,attachment,recipient_id,submitteddate,accountno,goal
	form.Add("recpimg", models.RecipientInfoMock.Image)
	form.Add("description", models.RecipientInfoMock.Description)
	form.Add("attachment", models.RecipientInfoMock.Attachment)
	form.Add("recipient",string(models.RecipientInfoMock.Recipient.RecipientNo))
	form.Add("date",models.RecipientInfoMock.Date.Format("2010-02-02 12:09:05"))
	form.Add("accoutNo", models.RecipientInfoMock.BankAccount.AccountNo)
	form.Add("goal",fmt.Sprintf("%f",models.RecipientInfoMock.Goal))



	resp, err := tc.PostForm(sURL+"/recipientInfo/create",form )
	fmt.Println(resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if bytes.Contains(body, []byte("mockyDesc")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestShowApproved(t *testing.T){
	tmpl := template.Must(template.ParseGlob("../../../web/template/*"))
	recipientInfoRepo := recipientInfo_repository.NewMockRecipientInfoRepository(nil)
	recipientInfoServ := recipientInfo_service.NewRecipientInfoService(recipientInfoRepo)

	recipientInfoHandler := NewRecipientInfoHandler (tmpl, recipientInfoServ,nil,)



	mux := http.NewServeMux()
	mux.HandleFunc("/category", recipientInfoHandler.ShowApproved)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL



	resp, err := tc.Get(sURL+"/category" )
	fmt.Println(resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("mockyDesc")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestSeeIndividual(t *testing.T){
	tmpl := template.Must(template.ParseGlob("../../../web/template/*"))
	recipientInfoRepo := recipientInfo_repository.NewMockRecipientInfoRepository(nil)
	recipientInfoServ := recipientInfo_service.NewRecipientInfoService(recipientInfoRepo)

	recipientInfoHandler := NewRecipientInfoHandler (tmpl, recipientInfoServ,nil,)



	mux := http.NewServeMux()

	mux.HandleFunc("/1" ,recipientInfoHandler.SeeIndividual)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL



	resp, err := tc.Get(sURL+"/1" )
	fmt.Println(resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("MockImg.jpg")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestViewRecipientSpecificInfo(t *testing.T){
	tmpl := template.Must(template.ParseGlob("../../../web/template/*"))
	recipientInfoRepo := recipientInfo_repository.NewMockRecipientInfoRepository(nil)
	recipientInfoServ := recipientInfo_service.NewRecipientInfoService(recipientInfoRepo)

	recipientInfoHandler := NewRecipientInfoHandler (tmpl, recipientInfoServ,nil,)



	mux := http.NewServeMux()
	mux.HandleFunc("/admin/specific", recipientInfoHandler.ViewSpecificRecipientInfo)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL
	resp, err := tc.Get(sURL+"/admin/specific?rNo=1" )
	fmt.Println(resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("mockyDesc")) {
		t.Errorf("want body to contain %q", body)
	}

}
func TestApproveRecipient(t *testing.T){
	tmpl := template.Must(template.ParseGlob("../../../web/template/*"))
	recipientInfoRepo := recipientInfo_repository.NewMockRecipientInfoRepository(nil)
	recipientInfoServ := recipientInfo_service.NewRecipientInfoService(recipientInfoRepo)

	recipientInfoHandler := NewRecipientInfoHandler (tmpl, recipientInfoServ,nil,)



	mux := http.NewServeMux()
	mux.HandleFunc("/admin/specific?rNo=1/approve", recipientInfoHandler.ApproveRecipient)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL
	form := url.Values{}
	form.Add("recipientInfoId", string(models.RecipientInfoMock.ID))
	resp, err := tc.PostForm(sURL+"/admin/specific?rNo=1/approve",form )
	fmt.Println(resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("mockyDesc")) {
		t.Errorf("want body to contain %q", body)
	}

}
func TestSelectMedicalCategory(t *testing.T){
	tmpl := template.Must(template.ParseGlob("../../../web/template/*"))
	catRepo := category_repository.NewMockCategoryRepository(nil)
	catServ := category_service.NewCategoryService(catRepo)
	categoryHandler := NewCategoryHandler (tmpl, catServ)

	mux := http.NewServeMux()

	mux.HandleFunc("/medical" ,categoryHandler.SelectMedicalCategory)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	resp, err := tc.Get(sURL+"/medical" )
	fmt.Println(resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("medical")) {
		t.Errorf("want body to contain %q", body)
	}
}
func TestSelectSpecificCategory(t *testing.T){
	tmpl := template.Must(template.ParseGlob("../../../web/template/*"))
	catRepo := category_repository.NewMockCategoryRepository(nil)
	catServ := category_service.NewCategoryService(catRepo)
	categoryHandler := NewCategoryHandler (tmpl, catServ)

	mux := http.NewServeMux()

	mux.HandleFunc("/specific" ,categoryHandler.SelectSpecificCategory)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	resp, err := tc.Get(sURL+"/specific" )
	fmt.Println(resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("mockcategory")) {
		t.Errorf("want body to contain %q", body)
	}
}


