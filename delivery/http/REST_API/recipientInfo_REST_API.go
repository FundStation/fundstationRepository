package REST_API

import (
	"encoding/json"
	"fmt"
	"github.com/FundStation/models"
	"github.com/FundStation/recipientInfo"
	"net/http"
	"path"
	"strconv"

	//"path"
	//"strconv"
)

type RecipientInfoApiHandler struct {
	recInfoApiService  recipientInfo.RecipientInfoService


}
func NewRecipientInfoApiHandler( ris recipientInfo.RecipientInfoService) *RecipientInfoApiHandler {
	return &RecipientInfoApiHandler{ recInfoApiService:ris}
}

func (riah *RecipientInfoApiHandler)GetRecipientsInfo(w http.ResponseWriter, r *http.Request){
	recipientsInfo,err := riah.recInfoApiService.SelectApproved()
	fmt.Println("recipients",recipientsInfo)
	if err != nil{
		return
	}
	output,err:= json.MarshalIndent(&recipientsInfo,"","\t\t")

	if err != nil{
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}
func (riah *RecipientInfoApiHandler)GetRecipientInfo(w http.ResponseWriter, r *http.Request){
	id,err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}
	recipientInfo,err := riah.recInfoApiService.ViewSpecificRecipientInfo(id)
	fmt.Println("recipient",recipientInfo)
	if err != nil{
		return
	}
	output,err:= json.MarshalIndent(&recipientInfo,"","\t\t")

	if err != nil{
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}

func (riah *RecipientInfoApiHandler)PostRecipientInfo(w http.ResponseWriter, r *http.Request){
	len := r.ContentLength
	body := make([]byte,len)
	r.Body.Read(body)
	var recipientInfo models.RecipientInfo
	json.Unmarshal(body,&recipientInfo)
	err := riah.recInfoApiService.CreateRecipientInfo(recipientInfo)
	if err != nil{
		fmt.Println("err",err)
		return
	}
	p := fmt.Sprintf("/v1/recipient/%d", recipientInfo.Recipient.RecipientNo)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}
func (riah *RecipientInfoApiHandler)PutRecipientInfo(w http.ResponseWriter, r *http.Request){
	id,err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}
	err=riah.recInfoApiService.ApproveRecipientInfo(id)
	if err != nil{
		return
	}
	w.WriteHeader(200)
	return
}
func (riah *RecipientInfoApiHandler)DeleteRecipientInfo(w http.ResponseWriter, r *http.Request){
	id,err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}
	err=riah.recInfoApiService.DeleteRecipientInfoById(id)
	if err != nil{
		fmt.Println("error occuredddd")
		return
	}
	w.WriteHeader(200)
	return
}
