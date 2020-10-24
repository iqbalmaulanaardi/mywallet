package repository

import (
	"github.com/iqbalmaulanaardi/mywallet/models"
)

func (s *Service) InsertSession(session models.Session) (result models.Session, err error) {
	err = s.DB.Raw(`insert into sessions values(?, ?, ?) ON CONFLICT ON CONSTRAINT sessions_user_id_key 
	DO UPDATE SET access_token = ? , is_active = true;`, session.UserID, session.AccessToken, session.IsActive, session.AccessToken).Scan(&result).Error
	return
}
func (s *Service) GetActiveSession(accessToken string) (result models.Session, err error) {
	err = s.DB.Raw(`select * from sessions where sessions.access_token = ? and sessions.is_active = ?;`, accessToken, true).Scan(&result).Error
	return
}
func (s *Service) Logout(userID uint64) (result models.Session, err error) {
	err = s.DB.Raw(`update sessions set is_active = false where sessions.user_id = ?;`, userID).Scan(&result).Error
	return
}
