package services

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mywallet/models"
	"github.com/iqbalmaulanaardi/mywallet/repository"
	"github.com/iqbalmaulanaardi/mywallet/requests"
	"github.com/iqbalmaulanaardi/mywallet/responses"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var (
		loginRequest requests.LoginRequest
		err          error
		user         models.User
		claims       responses.MyClaims
		signedToken  string
	)
	if err = c.Bind(&loginRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid Param(s)"})
		return
	}
	service := c.MustGet("service").(*repository.Service)
	if user, err = service.GetUserByUsername(loginRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect PIN!"})
		return
	}
	claims = responses.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "mywallet",
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		Username: user.Username,
		Email:    user.Email,
		UserID:   user.UserID,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	if signedToken, err = token.SignedString([]byte("9labqi6")); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	go service.TopUpBalance(user.UserID, 0.0)
	c.JSON(http.StatusOK, gin.H{"token": signedToken, "message": "Login Success!"})
	return
}
