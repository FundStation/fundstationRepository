package admin_repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/FundStation/models"
)

type PsqlAdminRepository struct {
	conn *sql.DB
}

func NewPsqlAdminRepository(Conn *sql.DB) *PsqlAdminRepository {
	return &PsqlAdminRepository{conn: Conn}
}
func (pr *PsqlAdminRepository) SelectAdmin(user string) (*models.Admin,error) {
	admin1 :=&models.Admin{}
	querystmt, err := pr.conn.Prepare("SELECT * FROM admins WHERE username=$1 ")

	if err != nil {
		return admin1,err
	}


	err = querystmt.QueryRow(user).Scan(&admin1.Username, &admin1.Password,&admin1.RoleID)


	if err != nil {
		return admin1,err
	}
	return admin1,nil

}
