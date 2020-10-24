package requests

import "errors"

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (r *RegisterRequest) Validate() (err error) {
	//validation
	if r.Username == "" || r.Password == "" || r.Email == "" {
		return errors.New("Invalid Param(s)")
	}
	return
}
