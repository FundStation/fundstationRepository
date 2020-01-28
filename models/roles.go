package models


// Role repesents application user roles
type Role struct {
	ID    uint
	Name  string `json:"type:varchar(255)"`
	Donor []Donor
}

type Session struct {
	ID         uint
	UUID       string `json:"type:varchar(255);not null"`
	Expires    int64  `json:"type:varchar(255);not null"`
	SigningKey []byte `json:"type:varchar(255);not null"`
}



	

