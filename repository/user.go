package repository

import (
	"github.com/iqbalmaulanaardi/mywallet/models"
	"github.com/iqbalmaulanaardi/mywallet/requests"
)

func (s *Service) Register(user requests.RegisterRequest) (result models.User, err error) {
	err = s.DB.Raw(`insert into users values(DEFAULT, ?, ?, ?);`, user.Username, user.Email, user.Password).Scan(&result).Error
	return
}
func (s *Service) GetUserByUsername(login requests.LoginRequest) (result models.User, err error) {
	err = s.DB.Raw(`select * from users where users.username = ?;`, login.Username).Scan(&result).Error
	return
}
