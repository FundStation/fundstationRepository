package recipientInfo_repository

import (
	"database/sql"
	"errors"

	"github.com/FundStation/models"
)

type PsqlRecipientInfoRepository struct {
	conn *sql.DB
}

func NewPsqlRecipientInfoRepository(Conn *sql.DB) *PsqlRecipientInfoRepository {
	return &PsqlRecipientInfoRepository{conn: Conn}
}

func (pr *PsqlRecipientInfoRepository) InsertRecipientInfo(ri models.RecipientInfo) error {

	err := pr.conn.QueryRow("INSERT INTO recipientinfo(image,description,attachment,recipient_id,submitteddate,accountno) VALUES($1, $2, $3, $4,$5,$6) returning id", ri.Image, ri.Description, ri.Attachment, ri.Recipient.RecipientNo, ri.Date, ri.BankAccount.AccountNo).Scan(&ri.ID)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PsqlRecipientInfoRepository) SelectRecipientInfo(recipientNo int) (recipientInfo models.RecipientInfo, err error) {
	var bankAccount models.BankAccount
	err = pr.conn.QueryRow("SELECT id, image, description,attachment, submitteddate,accountno FROM recipientinfo WHERE recipient_id=$1", recipientNo).Scan(&recipientInfo.ID, &recipientInfo.Image, &recipientInfo.Description, &recipientInfo.Attachment, &recipientInfo.Date, &bankAccount.AccountNo)
	recipientInfo.BankAccount = bankAccount
	if err != nil {
		return recipientInfo, errors.New("")
	}
	return recipientInfo, nil

}

func (pr *PsqlRecipientInfoRepository) UpdateRecipientInfo(recipientInfoNo int) (err error) {
	_,err = pr.conn.Exec("UPDATE recipientinfo SET approval = 'yes' WHERE id = $1",recipientInfoNo)
	if err != nil {
		return errors.New("Could not approve Recipient Info")
	}
	return nil

}
func (pr *PsqlRecipientInfoRepository) AccountExistsInfo(account string) bool {

	err := pr.conn.QueryRow("SELECT * FROM recipientinfo WHERE accountno=$1", account)

	if err != nil {
		return false
	}

	return true
}
func (pr *PsqlRecipientInfoRepository) SelectApproved() ( []models.RecipientInfo, error){
	recps := []models.RecipientInfo{}
	rinfo := models.RecipientInfo{}
	var bAct models.BankAccount
	selRec,err := pr.conn.Query("SELECT (image,description,submitteddate,accountno) FROM recipientinfo WHERE approval=$1", "yes")

	rinfo.BankAccount = bAct
	if err != nil {
		return recps,err
	}

	for selRec.Next() {
		err := selRec.Scan(&rinfo.Image,&rinfo.Description,bAct.AccountNo)
		if err != nil {
			return recps, errors.New("Couldnot")
		}
		recps = append(recps, rinfo)
	}
	return recps, nil
}

