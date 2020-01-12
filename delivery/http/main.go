package main

import (
	"database/sql"
	"github.com/FundStation/category/category_repository"
	"github.com/FundStation/category/category_service"
	"github.com/FundStation/donor/donor_repository"
	"github.com/FundStation/donor/donor_service"
	"github.com/FundStation/models"
	"github.com/FundStation/recipient/recipient_repository"
	"github.com/FundStation/recipient/recipient_service"
	"github.com/FundStation/recipientInfo/recipientInfo_repository"
	"github.com/FundStation/recipientInfo/recipientInfo_service"
	"github.com/FundStation/role/role_repository"
	"github.com/FundStation/role/role_service"
	"github.com/FundStation/tokens"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"net/http"
	"time"

	"github.com/FundStation/delivery/http/handler"
	"html/template"
)

var tmpl = template.Must(template.ParseGlob("web/template/*"))

func main() {


	csrfSignKey := []byte(tokens.GenerateRandomID(32))


	dbcon, err := sql.Open("postgres", "postgres://postgres:password@localhost/funddb?sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer dbcon.Close()

	donRepo := donor_repository.NewPsqlDonorRepository(dbcon)
	donServ := donor_service.NewDonorService(donRepo)

	recRepo := recipient_repository.NewPsqlRecipientRepository(dbcon)
	recServ := recipient_service.NewRecipientService(recRepo)

	sessionRepo := role_repository.NewSessionRepo(dbcon)
	sessionSrv := role_service.NewSessionService(sessionRepo)

	recInfoRepo := recipientInfo_repository.NewPsqlRecipientInfoRepository(dbcon)
	recInfoServ := recipientInfo_service.NewRecipientInfoService(recInfoRepo)

	catRepo := category_repository.NewPsqlCategoryRepository(dbcon)
	catServ := category_service.NewCategoryService(catRepo)


	roleRepo := role_repository.NewRoleRepository(dbcon)
	roleServ := role_service.NewRoleService(roleRepo)

	sess := configSess()
	dch:= handler.NewDonorHandler(tmpl, donServ, sessionSrv, roleServ, sess, csrfSignKey)
	rch:= handler.NewRecipientHandler(tmpl, recServ, sessionSrv, roleServ, sess, csrfSignKey)
	rich:= handler.NewRecipientInfoHandler(tmpl, recInfoServ, csrfSignKey)
	ch := handler.NewCategoryHandler(tmpl,catServ,csrfSignKey)


	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.Handle("/doners", dch.Authenticated(dch.Authorized(http.HandlerFunc(dch.Doners))))
	http.HandleFunc("/donor/signup", dch.DonorSignup)
	http.HandleFunc("/donor/login",dch.DonorLogin)
	http.HandleFunc("/donors",dch.Doners)

	http.HandleFunc("/recipient/signup",rch.RecipientSignup)
	http.HandleFunc("/recipient/login",rch.RecipientLogin)
	http.HandleFunc("/admin",rch.Recipients)

	http.HandleFunc("/recipientInfo",rich.CreateRecipientInfo)
	http.HandleFunc("/category",rich.SelectApproved)
	http.HandleFunc("/approve",rich.ApproveRecipient)
	http.HandleFunc("/admin/specific",rich.ViewSpecificRecipientInfo)


	http.Handle("/logout", dch.Authenticated(http.HandlerFunc(dch.Logout)))
	http.HandleFunc("/home",handler.Home)
	http.HandleFunc("/category/medical",ch.SelectMedicalCategory)
	http.HandleFunc("/category/women",ch.SelectWomenCategory)
	http.HandleFunc("/category/eduacational",ch.SelectEducationalCategory)
	http.HandleFunc("/category/orphanage",ch.SelectOrphanageCategory)
	http.HandleFunc("/category/religious",ch.SelectReligiousCategory)
	http.HandleFunc("/category/other",ch.SelectOtherCategory)
	http.HandleFunc("/category/specific",ch.SelectSpecificCategory)




	router := httprouter.New()
	router.GET("/v1/donor/:dlusername", dch.GetDonor)
	router.GET("/v1/donors", dch.GetDonors)
	router.POST("/v1/donor", dch.PostDonor)

	http.ListenAndServe(":8080",nil)
}

func configSess() *models.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := tokens.GenerateRandomID(32)
	signingString, err := tokens.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &models.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}
