package admin_repository

import (
	"database/sql"
	_ "github.com/lib/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type PsqlAdminRepository struct {
	conn *sql.DB
}

func NewPsqlAdminRepository(Conn *sql.DB) *PsqlAdminRepository {
	return &PsqlAdminRepository{conn: Conn}
}
