package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FundStation/donor"
	"github.com/FundStation/permission"
	"github.com/FundStation/role"
	"github.com/FundStation/session"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"strings"

	"github.com/FundStation/form"
	"github.com/FundStation/models"
	"github.com/FundStation/tokens"
	"html/template"
	"net/http"
	"net/url"
)


type DonorHandler struct {
	tmpl           *template.Template
	donorService    donor.DonorService
	sessionService role.SessionService
	donorSess       *models.Session
	loggedInDonor   *models.Donor
	donorRole       role.RoleService
	csrfSignKey    []byte
}
type contextKey string
var ctxUserSessionKey = contextKey("signed_in_user_session")


func NewDonorHandler(t *template.Template, ds donor.DonorService,
	sessServ role.SessionService, dRole role.RoleService,
	usrSess *models.Session, csKey []byte) *DonorHandler {
	return &DonorHandler{tmpl: t, donorService: ds, sessionService: sessServ,
		donorRole: dRole, donorSess: usrSess, csrfSignKey: csKey}
}

// Authenticated checks if a user is authenticated to access a given route
func (dch *DonorHandler) Authenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ok := dch.loggedIn(r)
		if !ok {
			http.Redirect(w, r, "/donor/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserSessionKey, dch.donorSess)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}


func (dch *DonorHandler) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if dch.loggedInDonor == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		role, errs := dch.donorRole.DonorRoles(dch.loggedInDonor)
		if errs != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}


			permitted := permission.HasPermission(r.URL.Path, role.Name, r.Method)
			if !permitted {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return

		}
		if r.Method == http.MethodPost {
			ok, err := tokens.ValidCSRF(r.FormValue("_csrf"), dch.csrfSignKey)
			if !ok || (err != nil) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}





func (dch *DonorHandler) Doners(w http.ResponseWriter, r *http.Request) {
	donors, errs :=dch.donorService.ViewAllDonor()
	if errs != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	token, err := tokens.CSRFToken(dch.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
		Donors     []models.Donor
		CSRF       string
	}{
		Values:     nil,
		VErrors:    nil,
		Donors:		donors,
		CSRF:       token,
	}
	fmt.Println(tmplData)
	dch.tmpl.ExecuteTemplate(w, "donors.html", tmplData)
}


func (dch *DonorHandler) DonorSignup(w http.ResponseWriter, r *http.Request) {
	token, err := tokens.CSRFToken(dch.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		newCatForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		dch.tmpl.ExecuteTemplate(w, "donorSignup.html", newCatForm)
	}

	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		newDonForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		newDonForm.Required("firstName", "lastName","add","userName","password","repassword","phone","email")
		newDonForm.MinLength("userName", 8)
		newDonForm.MinLength("phone",10)
		newDonForm.PasswordMatches("password","repassword")
		newDonForm.CSRF = token

		if !newDonForm.Valid() {
			dch.tmpl.ExecuteTemplate(w, "donorSignup.html", newDonForm)
			return
		}
		pExists := dch.donorService.PhoneExists(r.FormValue("phone"))
		if pExists {
			newDonForm.VErrors.Add("phone", "Phone Already Exists")
			dch.tmpl.ExecuteTemplate(w, "donorSignup.html", newDonForm)
			return
		}
		eExists := dch.donorService.EmailExists(r.FormValue("email"))
		if eExists {
			newDonForm.VErrors.Add("email", "Email Already Exists")
			dch.tmpl.ExecuteTemplate(w, "donorSignup.html", newDonForm)
			return
		}


		uExists := dch.donorService.UsernameExists(r.FormValue("userName"))
		if uExists {
			newDonForm.VErrors.Add("userName", "Username Already Exists")
			dch.tmpl.ExecuteTemplate(w, "donorSignup", newDonForm)
			return
		}


		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			newDonForm.VErrors.Add("password", "Password Could not be stored")
			dch.tmpl.ExecuteTemplate(w, "donorSignup.html",newDonForm)
			return
		}

		role, errs := dch.donorRole.RoleByName("DONOR")

		if  errs != nil {
			newDonForm.VErrors.Add("email", "could not assign role to the user")
			dch.tmpl.ExecuteTemplate(w, "donorSignup.html", newDonForm)
			fmt.Println(errs)
			return
		}


		donor := &models.Donor{
			FirstName:    r.FormValue("firstName"),
			LastName:     r.FormValue("lastName"),
			Address:      r.FormValue("add"),
			Occupation:   r.FormValue("occupation"),
			Username:     r.FormValue("userName"),
			Password:     string(hashedPassword),
			PhoneNumber:  r.FormValue("phone"),
			EmailAddress: r.FormValue("email"),
			RoleID:uint(role.ID),
			//BankAct: req.FormValue("bankAct"),
		}
		fmt.Println(donor)

		_,erro :=dch.donorService.SignupDonor(donor)
		if erro != nil{
			fmt.Println(err)

		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)

	}
}
func (dch *DonorHandler) DonorLogin(w http.ResponseWriter, r *http.Request) {
	token, err := tokens.CSRFToken(dch.csrfSignKey)
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
		dch.tmpl.ExecuteTemplate(w, "dlogin.html", loginForm)
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
		don, errs := dch.donorService.DonorByUsername(r.FormValue("dlusername"))
		if errs != nil{
			loginForm.VErrors.Add("generic", "Your username or password is wrong")
			dch.tmpl.ExecuteTemplate(w, "dlogin.html", loginForm)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(don.Password), []byte(r.FormValue("dlpassword")))

		if err == bcrypt.ErrMismatchedHashAndPassword {
			loginForm.VErrors.Add("generic", "Your username or password is wrong")
			dch.tmpl.ExecuteTemplate(w, "dlogin.html", loginForm)
			return
		}

		dch.loggedInDonor = don
		fmt.Println(err)

		claims := tokens.Claims(don.Username, dch.donorSess.Expires)
		session.Create(claims, dch.donorSess.UUID, dch.donorSess.SigningKey, w)
		newSess, errs := dch.sessionService.StoreSession(dch.donorSess)
		if errs != nil {
			loginForm.VErrors.Add("generic", "Failed to store session")
			dch.tmpl.ExecuteTemplate(w, "dlogin.html", loginForm)
			return
		}
		dch.donorSess = newSess
		roles, _ := dch.donorRole.DonorRoles(don)
		if dch.checkAdmin(roles) {
			http.Redirect(w, r, "/donor/signup", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)

	}
}

func (dch *DonorHandler) Logout(w http.ResponseWriter, r *http.Request) {
		donSess, _ := r.Context().Value(ctxUserSessionKey).(*models.Session)
		session.Remove(donSess.UUID, w)
		dch.sessionService.DeleteSession(donSess.UUID)
		http.Redirect(w, r, "/donor/login", http.StatusSeeOther)
	}

func (dch *DonorHandler) loggedIn(r *http.Request) bool {
	if dch.donorSess== nil {
		return false
	}
	donSess := dch.donorSess
	c, err := r.Cookie(donSess.UUID)
	if err != nil {
		return false
	}
	ok, err := session.Valid(c.Value, donSess.SigningKey)
	if !ok || (err != nil) {
		return false
	}
	return true
}

func (dch *DonorHandler) GetDonors(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	donors, errs := dch.donorService.ViewAllDonor();
	fmt.Println(donors)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(donors, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return

}

func (dch *DonorHandler) GetDonor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := ps.ByName("dlusername")
	//password:=ps.ByName("dlpassword")
	//donor:=models.Donor{
	//	Username :username,
	//	Password:password,
	//}
	//if err != nil {
	//	w.Header().Set("Content-Type", "application/json")
	//	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	//	return
	//}

	don,errs := dch.donorService.DonorByUsername(username)
	fmt.Println(don )

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(don, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
func (dch *DonorHandler) PostDonor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	donor:= &models.Donor{
		FirstName:r.FormValue("firstName"),
		LastName:r.FormValue("lastName"),
		Address:r.FormValue("add"),
		Occupation:r.FormValue("occupation"),
		Username:r.FormValue("userName"),
		Password:r.FormValue("password"),
		EmailAddress:r.FormValue("email"),
		PhoneNumber:r.FormValue("phone"),
	}

	err := json.Unmarshal(body, donor)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	donor, errs := dch.donorService.SignupDonor(donor)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/doner/%s", donor.Username)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

func (dch *DonorHandler) checkAdmin(r models.Role) bool {

		if strings.ToUpper(r.Name) == strings.ToUpper("donor") {
			return true
		}

	return false
}

