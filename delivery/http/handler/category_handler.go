package handler

import (
	"fmt"
	"github.com/FundStation2/category"
	"html/template"
	"net/http"
)

type CategoryHandler struct {
	tmpl           *template.Template
	categoryService    category.CategoryService
	csrfSignKey    []byte

}
func NewCategoryHandler(t *template.Template, cs category.CategoryService,csKey []byte) *CategoryHandler {
	return &CategoryHandler{tmpl: t, categoryService: cs,csrfSignKey:csKey}
}

func (cat *CategoryHandler) SelectMedicalCategory(w http.ResponseWriter, r *http.Request) {

	cato,err:=cat.categoryService.ViewCategory("medical")
	if err != nil{
		panic(err)
	}
	tmpl.ExecuteTemplate(w,"listCategory.layout",cato)
}
func (cat *CategoryHandler) SelectWomenCategory(w http.ResponseWriter, r *http.Request) {

	cato,err:=cat.categoryService.ViewCategory("women")
	if err != nil{
		panic(err)
	}
	tmpl.ExecuteTemplate(w,"listCategory.layout",cato)
}
func (cat *CategoryHandler) SelectEducationalCategory(w http.ResponseWriter, r *http.Request) {

	cato,err:=cat.categoryService.ViewCategory("education")
	if err != nil{
		panic(err)
	}
	tmpl.ExecuteTemplate(w,"listCategory.layout",cato)
}
func (cat *CategoryHandler) SelectOrphanageCategory(w http.ResponseWriter, r *http.Request) {

	cato,err:=cat.categoryService.ViewCategory("orphanage")
	if err != nil{
		panic(err)
	}
	tmpl.ExecuteTemplate(w,"listCategory.layout",cato)
}

func (cat *CategoryHandler) SelectReligiousCategory(w http.ResponseWriter, r *http.Request) {

	cato,err:=cat.categoryService.ViewCategory("religious")
	if err != nil{
		panic(err)
	}
	tmpl.ExecuteTemplate(w,"listCategory.layout",cato)
}

func (cat *CategoryHandler) SelectOtherCategory(w http.ResponseWriter, r *http.Request) {

	cato,err:=cat.categoryService.ViewCategory("other")
	if err != nil{
		panic(err)
	}
	tmpl.ExecuteTemplate(w,"listCategory.layout",cato)
}
func (cat *CategoryHandler) SelectSpecificCategory(w http.ResponseWriter,r *http.Request){
	name:=r.FormValue("catName")
	fmt.Println(name)
	cato,err:=cat.categoryService.ViewSpecificCategory(name)
	if err != nil{
		fmt.Println(err)
	}
	tmpl.ExecuteTemplate(w,"specificEntity.layout",cato)
}





