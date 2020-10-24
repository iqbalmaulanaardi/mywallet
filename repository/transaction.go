package repository

import (
	"github.com/iqbalmaulanaardi/mywallet/models"
	"github.com/iqbalmaulanaardi/mywallet/responses"
)

func (s *Service) InsertTransaction(transfer models.Transaction) (result models.Transaction, err error) {
	err = s.DB.Model(models.Transaction{}).Create(&transfer).Error
	return
}
func (s *Service) GetTransactionsByUserID(userID uint64) (result []responses.TransactionListResponse, err error) {
	err = s.DB.Raw(`select * from (select '-' as source_user ,users.username as dest_user, 'DEBIT' as type,transactions.amount, transactions.created_at from transactions inner join users on transactions.to = users.user_id where transactions.from = ? UNION select users.username as source_user, '-' as dest_user, 'CREDIT' as type,transactions.amount, transactions.created_at from transactions inner join users on transactions.from = users.user_id where transactions.to = ?) results order by created_at DESC;`, userID, userID).Find(&result).Error
	return
}
