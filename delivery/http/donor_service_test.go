package main

import (
	"bytes"
	"github.com/FundStation/delivery/http/handler"
	"github.com/FundStation/donor/donor_repository"
	"github.com/FundStation/donor/donor_service"
	"github.com/FundStation/models"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestDonor(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../web/template/*"))

	donorRepo := donor_repository.NewMockDonorRepo(nil)
	donorServ := donor_service.NewDonorService(donorRepo)

	donorHandler := handler.NewDonorHandler (tmpl, donorServ,nil,nil,nil,nil,)


	mux := http.NewServeMux()
	mux.HandleFunc("/donors",donorHandler.Doners)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/donors")
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

func TestDonorNew(t *testing.T) {
	tmpl := template.Must(template.ParseGlob("../../web/template/*"))
	donorRepo := donor_repository.NewMockDonorRepo(nil)
	donorServ := donor_service.NewDonorService(donorRepo)

	donorHandler := handler.NewDonorHandler (tmpl, donorServ,nil,nil,nil,nil)



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

