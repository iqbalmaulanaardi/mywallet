package requests

import "errors"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginRequest) Validate() (err error) {
	//validation
	if l.Username == "" || l.Password == "" {
		return errors.New("Invalid Param(s)")
	}
	return
}
