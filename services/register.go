package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mywallet/repository"
	"github.com/iqbalmaulanaardi/mywallet/requests"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var (
		registerRequest requests.RegisterRequest
		err             error
	)
	if err = c.Bind(&registerRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid param(s)"})
		return
	}
	service := c.MustGet("service").(*repository.Service)
	EncryptPassword(&registerRequest)
	if _, err = service.Register(registerRequest); err != nil && err.Error() != "record not found" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registered. Please Login!"})
	return
}
func EncryptPassword(registerRequest *requests.RegisterRequest) {
	var (
		err  error
		hash []byte
	)
	if hash, err = bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.MinCost); err != nil {
		return
	}
	registerRequest.Password = string(hash)
	return
}
