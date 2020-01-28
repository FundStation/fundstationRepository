package models

import "time"

type RecipientInfo struct {
	ID          int               `json:"id"`
	Image       string			  `json:"image"`
	Description string			  `json:"description"`
	Attachment  string			  `json:"attachment"`
	Recipient   *Recipient		  `json:"recipient_id"`
	Date        time.Time		  `json:"submitteddate"`
	BankAccount BankAccount       `json:"account"`
	Goal        float64			  `json:"goal"`
}
