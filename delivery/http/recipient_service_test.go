package main

import (
	"bytes"
	"github.com/FundStation2/delivery/http/handler"
	"github.com/FundStation2/models"
	"github.com/FundStation2/recipient/recipient_repository"
	"github.com/FundStation2/recipient/recipient_service"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRecipient(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../web/template/*"))

	recRepo := recipient_repository.NewMockRecipientRepo(nil)
	recServ := recipient_service.NewRecipientService(recRepo)

	recHandler := handler.NewRecipientHandler (tmpl, recServ,nil,nil,nil,nil,)


	mux := http.NewServeMux()
	mux.HandleFunc("/admin",recHandler.Recipients)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin")
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

func TestRecipientNew(t *testing.T) {
	tmpl := template.Must(template.ParseGlob("../../web/template/*"))
	recRepo := recipient_repository.NewMockRecipientRepo(nil)
	recServ := recipient_service.NewRecipientService(recRepo)

	recHandler := handler.NewRecipientHandler (tmpl, recServ,nil,nil,nil,nil,)


	mux := http.NewServeMux()
	mux.HandleFunc("/recipient/signup", recHandler.RecipientSignup)
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
	form.Add("repassword",models.RecipientMock.Password)
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

