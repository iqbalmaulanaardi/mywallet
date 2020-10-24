package requests

import "errors"

type TransferRequest struct {
	DestUser string  `json:"dest_user"`
	Amount   float64 `json:"amount"`
}

func (t *TransferRequest) Validate() (err error) {
	//validation
	if t.DestUser == "" {
		return errors.New("Invalid Param(s)")
	}
	if t.Amount <= 0 {
		return errors.New("Amount must be larger than 0")
	}
	return
}
