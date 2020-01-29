package main

import (
	"database/sql"
	"github.com/FundStation/admin/admin_repository"
	"github.com/FundStation/admin/admin_service"
	"github.com/FundStation/category/category_repository"
	"github.com/FundStation/category/category_service"
	"github.com/FundStation/delivery/http/REST_API"
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
	"net/http"
	"time"

	"github.com/FundStation/delivery/http/handler"
	"html/template"
)



func main() {
	var tmpl = template.Must(template.ParseGlob("web/template/*"))

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

	admRepo := admin_repository.NewPsqlAdminRepository(dbcon)
	admServ := admin_service.NewAdminService(admRepo)
	
	disRepo := disaster_repository.NewPsqlDisasterRepository(dbcon)
	disServ := disaster_service.NewDisasterService(disRepo)


	roleRepo := role_repository.NewRoleRepository(dbcon)
	roleServ := role_service.NewRoleService(roleRepo)


	sess := configSess()
	hch := handler.NewHomeHandler(tmpl)
	dch:= handler.NewDonorHandler(tmpl, donServ, sessionSrv, roleServ, sess, csrfSignKey)
	rch:= handler.NewRecipientHandler(tmpl, recServ, sessionSrv, roleServ, sess, csrfSignKey)
	rich:= handler.NewRecipientInfoHandler(tmpl, recInfoServ, csrfSignKey)
	ch := handler.NewCategoryHandler(tmpl,catServ)
	ah := handler.NewAdminHandler(tmpl,admServ,sessionSrv,roleServ,sess,csrfSignKey)
	dish := handler.NewDisasterHandler(tmpl,disServ)


	dah:= REST_API.NewDonorApiHandler(donServ)
	rah:= REST_API.NewRecipientApiHandler(recServ)
	riah:=REST_API.NewRecipientInfoApiHandler(recInfoServ)

	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.Handle("/donors", dch.Authenticated(dch.Authorized(http.HandlerFunc(dch.Doners))))
	http.HandleFunc("/donor", dch.Donor)
	http.HandleFunc("/donor/signup",dch.DonorSignup)
	//http.Handle("/donor/signup",dch.Authenticated(dch.Authorized(http.HandlerFunc(dch.DonorSignup))))
	//http.HandleFunc("/donor/signup", dch.DonorSignup)
	http.HandleFunc("/donor/login",dch.DonorLogin)
	//http.HandleFunc("/donors",dch.Doners)
	
	http.HandleFunc("/disaster",dish.SelectDisasters)
	http.HandleFunc("/event",dish.SelectEvents)
	
	http.HandleFunc("/recipient/signup",rch.RecipientSignup)
	http.HandleFunc("/recipient/login",rch.RecipientLogin)
	http.HandleFunc("/admin",ah.AdminLogin)
	http.Handle("/admin/recipients",ah.Authenticated(ah.Authorized(http.HandlerFunc(rch.Recipients))))
	http.Handle("/logout",rch.Authenticated(http.HandlerFunc(rch.Logout)))

	http.HandleFunc("/recipientInfo",rich.RecipientInfo)                           //Check this later
	http.HandleFunc("/recipientInfo/create",rich.CreateRecipientInfo)
	http.HandleFunc("/category",rich.ShowApproved)

	http.HandleFunc("/individual/all",rich.SeeAll)
	gd:="/"+handler.Path
	http.HandleFunc(gd,rich.SeeIndividual)
	http.HandleFunc("/approve",rich.ApproveRecipient)
	http.HandleFunc("/admin/specific",rich.ViewSpecificRecipientInfo)


//	http.Handle("/logout", dch.Authenticated(http.HandlerFunc(dch.Logout)))
	http.HandleFunc("/home",hch.Home)
	http.HandleFunc("/category/medical",ch.SelectMedicalCategory)
	http.HandleFunc("/category/women",ch.SelectWomenCategory)
	http.HandleFunc("/category/eduacation",ch.SelectEducationalCategory)
	http.HandleFunc("/category/orphanage",ch.SelectOrphanageCategory)
	http.HandleFunc("/category/religious",ch.SelectReligiousCategory)
	http.HandleFunc("/category/other",ch.SelectOtherCategory)
	http.HandleFunc("/category/specific",ch.SelectSpecificCategory)
//	http.Handle("category/donation",dch.Authenticated(dch.Authorized(http.HandlerFunc(ch.Donation))))
	http.HandleFunc("/category/donation",ch.Donation)
	http.HandleFunc("/money/transfer", ch.DonateMoney)


	//for the REST API
	http.HandleFunc("/v1/donors",dah.GetDonors)
	http.HandleFunc("/v1/donor/",dah.GetDonor)
	http.HandleFunc("/v1/donor",dah.PostDonor)
	http.HandleFunc("/v1/donor/update/",dah.PutDonor)
	http.HandleFunc("/v1/donor/delete/",dah.DeleteDonor)

	http.HandleFunc("/v1/recipients",rah.GetRecipients)
	http.HandleFunc("/v1/recipient/",rah.GetRecipient)
	http.HandleFunc("/v1/recipient",rah.PostRecipient)
	http.HandleFunc("/v1/recipient/update/",rah.PutRecipient)
	http.HandleFunc("/v1/recipient/delete/",rah.DeleteRecipient)

	http.HandleFunc("/v1/recipientsinfo",riah.GetRecipientsInfo)
	http.HandleFunc("/v1/recipientinfo/",riah.GetRecipientInfo)
	http.HandleFunc("/v1/recipientinfo",riah.PostRecipientInfo)
	http.HandleFunc("/v1/recipientinfo/update/",riah.PutRecipientInfo)
	http.HandleFunc("/v1/recipientinfo/delete/",riah.DeleteRecipientInfo)



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
