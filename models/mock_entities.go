package models

import "time"

// RoleMock mocks user role entity
var RoleMock = Role{
	ID:    1,
	Name:  "Mock Role 01",
	Donor: []Donor{},
}



// SessionMock mocks sessions of loged in user
var SessionMock = Session{
	ID:         1,
	UUID:       "_session_one",
	SigningKey: []byte("FundStationApp"),
	Expires:    0,
}


var DonorMock = Donor{
	DonorNo:      1,
	FirstName:    "Mock donor 01",
	LastName:     "Mock donor1 01",
	Address:      "mockAdd",
	Username:     "mockyuser",
	Password:     "password",
	PhoneNumber:  "0911111111",
	EmailAddress: "mockuser@gmail.com",
}

var BankMock = BankAccount{
	AccountNo:"1111111111111",
	CurrentBalance:0,
}

var RecipientMock = Recipient{
	RecipientNo:      1,
	FirstName:    "Mock recipient 01",
	LastName:     "Mock recipient1 01",
	Address:      "mockAdd",
	Username:     "mockyuser",
	Password:     "password",
	PhoneNumber:  "0911111111",
	EmailAddress: "mockuser@gmail.com",
}
var RecipientInfoMock = RecipientInfo{
	ID:1,
	Image:"MockImg.jpg",
	Description:"mockyDesc",
	Attachment:"mockAttac",
	Recipient:&RecipientMock,
	Date:time.Time{},
	BankAccount:BankMock,
	Goal:0,

}

var DonationInfoMock = DonationInfo{
	RecipientNo:RecipientMock.RecipientNo,
	FirstName:RecipientMock.FirstName,
	LastName:RecipientMock.LastName,
	Address:RecipientMock.Address,
	PhoneNumber:RecipientMock.PhoneNumber,
	EmailAddress:RecipientMock.EmailAddress,
	Image:RecipientInfoMock.Image,
	Description:RecipientInfoMock.Description,
	Attachment:RecipientInfoMock.Attachment,
	Date:time.Time{},
	AccountNo:RecipientInfoMock.BankAccount.AccountNo,
	Goal:RecipientInfoMock.Goal,
	Balance:RecipientInfoMock.BankAccount.CurrentBalance,
}

var CategoryMock = Category{
	Name:"mockcategory",
	CatType:"mocktype",
	Description:"mockyDesc",
	Image:"MockImg.jpg",
}
