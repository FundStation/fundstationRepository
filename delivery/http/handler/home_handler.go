package handler

import (
	"html/template"
	"net/http"
)


var tmpl = template.Must(template.ParseGlob("web/template/*"))

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w,"home.layout",nil)
}

