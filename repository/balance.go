package repository

import "github.com/iqbalmaulanaardi/mywallet/models"

func (s *Service) TopUpBalance(userID uint64, amount float64) (balance models.Balance, err error) {
	err = s.DB.Raw(`insert into balances values(?, ?);`, userID, amount).Scan(&balance).Error
	return
}
func (s *Service) GetBalance(userID uint64) (balance models.Balance, err error) {
	err = s.DB.Raw(`select * from balances where balances.user_id = ?;`, userID).Scan(&balance).Error
	return
}
