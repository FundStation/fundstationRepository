package models

import "time"

type DonationInfo struct {
	RecipientNo         int
	FirstName    		string
	LastName     		string
	Address      		string
	PhoneNumber  		string
	EmailAddress 		string
	Image       		string
	Description			string
	Attachment 			string
	Date        		time.Time
	AccountNo 			string
	Goal                float64
	Balance             float64
	Current             float64
}

