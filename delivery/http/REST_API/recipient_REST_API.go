package REST_API

import (
	"encoding/json"
	"fmt"
	"github.com/FundStation/models"
	"github.com/FundStation/recipient"
	"net/http"
	"path"
	"strconv"

	//"path"
	//"strconv"
)

type RecipientApiHandler struct {
	recApiService  recipient.RecipientService


}
func NewRecipientApiHandler( rs recipient.RecipientService) *RecipientApiHandler {
	return &RecipientApiHandler{ recApiService:rs}
}

func (rah *RecipientApiHandler)GetRecipients(w http.ResponseWriter, r *http.Request){
	recipients,err := rah.recApiService.ViewAllRecipient()
	fmt.Println("recipients",recipients)
	if err != nil{
		return
	}
	output,err:= json.MarshalIndent(&recipients,"","\t\t")

	if err != nil{
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}
func (rah *RecipientApiHandler)GetRecipient(w http.ResponseWriter, r *http.Request){
	id,err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}
	recipient,err := rah.recApiService.RecipientById(id)
	fmt.Println("recipient",recipient)
	if err != nil{
		return
	}
	output,err:= json.MarshalIndent(&recipient,"","\t\t")

	if err != nil{
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}

func (rah *RecipientApiHandler)PostRecipient(w http.ResponseWriter, r *http.Request){
	len := r.ContentLength
	body := make([]byte,len)
	r.Body.Read(body)
	var recipient models.Recipient
	json.Unmarshal(body,&recipient)
	_,err := rah.recApiService.SignupRecipient(&recipient)
	if err != nil{
		fmt.Println("err",err)
		return
	}
	p := fmt.Sprintf("/v1/recipient/%d", recipient.RecipientNo)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}
func (rah *RecipientApiHandler)PutRecipient(w http.ResponseWriter, r *http.Request){
	id,err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}
	recipient,err:=rah.recApiService.RecipientById(id)
	if err != nil{
		return
	}
	len := r.ContentLength
	body := make([]byte,len)
	r.Body.Read(body)
	json.Unmarshal(body,&recipient)
	err = rah.recApiService.UpdateRecipientById(recipient)
	if err != nil{
		return
	}
	w.WriteHeader(200)
	return
}
func (rah *RecipientApiHandler)DeleteRecipient(w http.ResponseWriter, r *http.Request){
	id,err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil{
		return
	}
	recipient,err:=rah.recApiService.RecipientById(id)
	if err != nil{
		return
	}
	err = rah.recApiService.DeleteRecipientById(recipient)
	if err != nil{
		return
	}
	w.WriteHeader(200)
	return
}
