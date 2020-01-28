package handler

import (
	"context"
	"fmt"
	"github.com/FundStation2/admin"
	"github.com/FundStation2/form"
	"github.com/FundStation2/models"
	"github.com/FundStation2/permission"
	"github.com/FundStation2/role"
	"github.com/FundStation2/session"
	"github.com/FundStation2/tokens"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)


type AdminHandler struct {
	tmpl           *template.Template
	adminService    admin.AdminService
	sessionService role.SessionService
	adminSess       *models.Session
	loggedInAdmin   *models.Admin
	adminRole       role.RoleService
	csrfSignKey    []byte

}




func NewAdminHandler(t *template.Template, as admin.AdminService,
	sessServ role.SessionService, aRole role.RoleService,
	usrSess *models.Session, csKey []byte) *AdminHandler {
	return &AdminHandler{tmpl: t, adminService: as, sessionService: sessServ,
		adminRole: aRole, adminSess: usrSess, csrfSignKey: csKey}
}
//func NewAdminHandler(t *template.Template, as admin.AdminService) *AdminHandler {
//	return &AdminHandler{tmpl: t, adminService: as,}
//}
func (ach *AdminHandler) Authenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ok := ach.loggedIn(r)
		if !ok {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			fmt.Println("not")
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserSessionKey,ach.adminSess)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func (ach *AdminHandler) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ach.loggedInAdmin == nil {
			fmt.Println("s2")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		role, errs := ach.adminRole.AdminRoles(ach.loggedInAdmin)
		if errs != nil {
			fmt.Println(errs)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}


		permitted := permission.HasPermission(r.URL.Path, role.Name, r.Method)
		if !permitted {
			fmt.Println("s1")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return

		}
		if r.Method == http.MethodPost {
			ok, err := tokens.ValidCSRF(r.FormValue("_csrf"), ach.csrfSignKey)

			if !ok || (err != nil) {
				fmt.Println("s2")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (ach *AdminHandler) AdminLogin(w http.ResponseWriter, r *http.Request) {
	token, err := tokens.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		loginForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		ach.tmpl.ExecuteTemplate(w, "adminLogin.html", loginForm)
		return
	}
	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		loginForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		adm, errs := ach.adminService.LoginAdmin(r.FormValue("Ausername"))
		fmt.Println("don",errs)
		if errs != nil{
			loginForm.VErrors.Add("generic", "Username or Password is wrong")
			ach.tmpl.ExecuteTemplate(w, "adminLogin.html", loginForm)
			return
		}

		if adm.Password != r.FormValue("Apassword"){
			loginForm.VErrors.Add("generic", "Username or Password is wrong")
			ach.tmpl.ExecuteTemplate(w, "adminLogin.html", loginForm)
			return
		}

		ach.loggedInAdmin = adm
		fmt.Println(err)

		claims := tokens.Claims(adm.Username,ach.adminSess.Expires)
		session.Create(claims, ach.adminSess.UUID, ach.adminSess.SigningKey, w)
		newSess, errs := ach.sessionService.StoreSession(ach.adminSess)
		if errs != nil {
			loginForm.VErrors.Add("generic", "Failed to store session")
			ach.tmpl.ExecuteTemplate(w, "adminLogin.html", loginForm)
			return
		}
		ach.adminSess= newSess
		roles, _ := ach.adminRole.AdminRoles(adm)
		if ach.checkAdmin(roles) {
			http.Redirect(w, r, "/admin/recipients", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/admin", http.StatusSeeOther)

	}
}

func (ach *AdminHandler) loggedIn(r *http.Request) bool {
	if ach.adminSess== nil {
		return false
	}
	adminSess := ach.adminSess
	c, err := r.Cookie(adminSess.UUID)
	if err != nil {
		return false
	}
	ok, err := session.Valid(c.Value, adminSess.SigningKey)
	if !ok || (err != nil) {
		return false
	}
	return true
}
func (ach *AdminHandler) checkAdmin(r models.Role) bool {

	if strings.ToUpper(r.Name) == strings.ToUpper("admin") {
		return true
	}

	return false
}
//func (ach *AdminHandler) AdminLogin(res http.ResponseWriter, req *http.Request) {
//	if req.Method == http.MethodPost {
//		form := models.Admin{
//			Username: req.FormValue("Ausername"),
//			Password: req.FormValue("Apassword"),
//			//BankAct: req.FormValue("bankAct"),
//		}
//		err := ach.adminService.LoginAdmin(form)
//
//		if err != nil {
//			fmt.Print(err)
//		}
//		http.Redirect(res, req, "/admin", http.StatusSeeOther)
//	} else {
//
//		ach.tmpl.ExecuteTemplate(res, "adminLogin.html", nil)
//
//	}
//}