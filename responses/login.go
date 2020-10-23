package responses

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email    string `json:"email"`
	UserID   uint64 `json:"user_id`
}
