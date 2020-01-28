package REST_API

import (
	"encoding/json"
	"fmt"
	"github.com/FundStation/donor"
	"github.com/FundStation/models"
	"net/http"
	"path"
	"strconv"

	//"path"
	//"strconv"
)

type DonorApiHandler struct {
	donorApiService  donor.DonorService


}
func NewDonorApiHandler( ds donor.DonorService) *DonorApiHandler {
	return &DonorApiHandler{ donorApiService: ds}
}

func (dah *DonorApiHandler)GetDonors(w http.ResponseWriter, r *http.Request){
	//id,err := strconv.Atoi(path.Base(r.URL.Path))
	//if err != nil{
	//	return
	//}
	donors,err := dah.donorApiService.ViewAllDonor()
	fmt.Println("donorss",donors)
	if err != nil{
		return
	}
	output,err:= json.MarshalIndent(&donors,"","\t\t")

	if err != nil{
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}
func (dah *DonorApiHandler)GetDonor(w http.ResponseWriter, r *http.Request){
	id,err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}
	donor,err := dah.donorApiService.DonorById(id)
	fmt.Println("donorss",donor)
	if err != nil{
		return
	}
	output,err:= json.MarshalIndent(&donor,"","\t\t")

	if err != nil{
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}

func (dah *DonorApiHandler)PostDonor(w http.ResponseWriter, r *http.Request){
	len := r.ContentLength
	body := make([]byte,len)
	r.Body.Read(body)
	var donor models.Donor
	//rId,_:=strconv.Atoi(r.FormValue("role_id"))
	json.Unmarshal(body,&donor)
	_,err := dah.donorApiService.SignupDonor(&donor)
	if err != nil{
		fmt.Println("err",err)
		return
	}
	fmt.Println(donor)
	p := fmt.Sprintf("/v1/donor/%d", donor.DonorNo)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}
func (dah *DonorApiHandler)PutDonor(w http.ResponseWriter, r *http.Request){
	id,err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}
	donor,err:=dah.donorApiService.DonorById(id)
	if err != nil{
		return
	}
	len := r.ContentLength
	body := make([]byte,len)
	r.Body.Read(body)
	json.Unmarshal(body,&donor)
	err = dah.donorApiService.UpdateDonorById(donor)
	if err != nil{
		return
	}
	w.WriteHeader(200)
	return
}
func (dah *DonorApiHandler)DeleteDonor(w http.ResponseWriter, r *http.Request){
	id,err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}
	donor,err:=dah.donorApiService.DonorById(id)
	if err != nil{
		return
	}
	err = dah.donorApiService.DeleteDonorById(donor)
	if err != nil{
		return
	}
	w.WriteHeader(200)
	return
}
