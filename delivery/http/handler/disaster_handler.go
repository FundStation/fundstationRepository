package handler

import (
	"github.com/FundStation2/disaster"
	"html/template"
	"net/http"
)



type DisasterHandler struct {
	tmpl           *template.Template
	disasterService    disaster.DisasterService


}

func NewDisasterHandler(t *template.Template, cs disaster.DisasterService) *DisasterHandler{
	return &DisasterHandler{tmpl: t, disasterService: cs,}
}

func (cat *DisasterHandler) SelectDisasters(w http.ResponseWriter, r *http.Request) {

	dis,err:=cat.disasterService.ViewDisaster()
	if err != nil{
		panic(err)
	}
	cat.tmpl.ExecuteTemplate(w,"disaster.layout",dis)
}
func (cat *DisasterHandler) SelectEvents(w http.ResponseWriter, r *http.Request) {

	eve,err:=cat.disasterService.ViewEvent()
	if err != nil{
		panic(err)
	}
	cat.tmpl.ExecuteTemplate(w,"event.layout",eve)
}