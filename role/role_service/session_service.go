package role_service

import (
	"github.com/FundStation2/models"
	"github.com/FundStation2/role"
)


type SessionServiceImpl struct {
	sessionRepo role.SessionRepository
}


func NewSessionService(sessRepository role.SessionRepository) role.SessionService {
	return &SessionServiceImpl{sessionRepo: sessRepository}
}


func (ss *SessionServiceImpl) Session(sessionID string) (*models.Session, error) {
	sess, errs := ss.sessionRepo.Session(sessionID)
	if errs != nil {
		return nil, errs
	}
	return sess, errs
}


func (ss *SessionServiceImpl) StoreSession(session *models.Session) (*models.Session, error) {
	sess, errs := ss.sessionRepo.StoreSession(session)
	if errs != nil  {
		return nil, errs
	}
	return sess, errs
}


func (ss *SessionServiceImpl) DeleteSession(sessionID string) (*models.Session, error) {
	sess, errs := ss.sessionRepo.DeleteSession(sessionID)
	if errs != nil {
		return nil, errs
	}
	return sess, errs
}
