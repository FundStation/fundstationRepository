package role_repository

import (
	"database/sql"
	"github.com/FundStation2/models"
	"github.com/FundStation2/role"
)


type SessionRepo struct {
	conn *sql.DB
}


func NewSessionRepo(db *sql.DB)(role.SessionRepository){
	return &SessionRepo{conn: db}
}


func (sr *SessionRepo) Session(sessionID string) (*models.Session, error) {
	session := &models.Session{}


	querystmt, err := sr.conn.Prepare("SELECT * FROM session WHERE uuid=$1")

	if err != nil {
		return session, err
	}

	err = querystmt.QueryRow(sessionID).Scan(session.ID, session.UUID,session.Expires,session.SigningKey)
	if err != nil {
		return session, err
	}

	return session, nil

}

// StoreSession stores a given session
func (sr *SessionRepo) StoreSession(session *models.Session) (*models.Session, error) {
	sess := session
	err := sr.conn.QueryRow("INSERT INTO session(uuid,expires,signingkey) VALUES($1, $2, $3) returning id", session.UUID, session.Expires, session.SigningKey).Scan(&sess.ID)

	if err != nil{
		return sess, err
	}
	return sess, nil



}

// DeleteSession deletes a given session
func (sr *SessionRepo) DeleteSession(sessionID string) (*models.Session, error) {
	sess, err := sr.Session(sessionID)
	if err != nil{
		return sess,err
	}

	_,err= sr.conn.Exec("DELETE FROM session WHERE uuid=$1",sessionID)


	if err != nil {
		return sess, err
	}



	return sess, nil

}
