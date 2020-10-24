package services

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mywallet/models"
	"github.com/iqbalmaulanaardi/mywallet/repository"
)

func GetBalance(c *gin.Context) {
	var (
		err                 error
		userInfo            map[string]interface{}
		authorizationHeader string
		balance             models.Balance
	)
	authorizationHeader = c.Request.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	if userInfo, err = ValidateSession(c, authorizationHeader); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	userIDStr := fmt.Sprintf("%v", userInfo["UserID"])
	userID, _ := strconv.ParseInt(userIDStr, 16, 16)
	service := c.MustGet("service").(*repository.Service)
	if balance, err = service.GetBalance(uint64(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success Get Balance", "balance": balance.Balance})
	return
}
