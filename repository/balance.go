package repository

import "github.com/iqbalmaulanaardi/mywallet/models"

func (s *Service) InitBalance(userID uint64, amount float64) (balance models.Balance, err error) {
	err = s.DB.Raw(`insert into balances values(?, ?) ON CONFLICT ON CONSTRAINT balances_user_id_key 
	DO NOTHING;`, userID, amount).Scan(&balance).Error
	return
}
func (s *Service) GetBalance(userID uint64) (balance models.Balance, err error) {
	err = s.DB.Raw(`select * from balances where balances.user_id = ?;`, userID).Scan(&balance).Error
	return
}
func (s *Service) TopUpBalance(destUserID uint64, amount float64) (balance models.Balance, err error) {
	err = s.DB.Raw(`UPDATE balances SET balance=(select balance from balances where balances.user_id = ?) + ? WHERE balances.user_id = ?;`, destUserID, amount, destUserID).Scan(&balance).Error
	return
}
func (s *Service) DeductBalance(destUserID uint64, amount float64) (balance models.Balance, err error) {
	err = s.DB.Raw(`UPDATE balances SET balance=(select balance from balances where balances.user_id = ?) - ? WHERE balances.user_id = ?;`, destUserID, amount, destUserID).Scan(&balance).Error
	return
}
