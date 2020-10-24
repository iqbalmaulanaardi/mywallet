package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Param(s)"})
		return
	}
	if err = loginRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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
	go service.InsertSession(models.Session{
		UserID:      user.UserID,
		AccessToken: signedToken,
		IsActive:    true,
	})
	go service.InitBalance(user.UserID, 0.0)
	c.JSON(http.StatusOK, gin.H{"token": signedToken, "message": "Login Success!"})
	return
}

func ValidateSession(c *gin.Context, authorizationHeader string) (userInfo map[string]interface{}, err error) {
	var (
		claims      responses.MyClaims
		tokenString string
	)
	tokenString = strings.Replace(authorizationHeader, "Bearer ", "", -1)
	service := c.MustGet("service").(*repository.Service)
	if _, err = service.GetActiveSession(tokenString); err != nil {
		err = errors.New("Invalid Session, please relogin")
		return
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	if token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return []byte("9labqi6"), nil
	}); err != nil {
		return
	}
	claims2, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return
	}
	ctx := context.WithValue(context.Background(), "user_info", claims2)
	userInfo = ctx.Value("user_info").(jwt.MapClaims)
	return
}
