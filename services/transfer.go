package services

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mywallet/models"
	"github.com/iqbalmaulanaardi/mywallet/repository"
	"github.com/iqbalmaulanaardi/mywallet/requests"
)

func Transfer(c *gin.Context) {
	var (
		err                 error
		userInfo            map[string]interface{}
		authorizationHeader string
		transferRequest     requests.TransferRequest
		balance             models.Balance
		destUser            models.User
	)
	if err = c.Bind(&transferRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Param(s)"})
		return
	}
	if err = transferRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
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
	//validate dest_user
	if destUser, err = service.GetUserByUsername(requests.LoginRequest{
		Username: transferRequest.DestUser,
	}); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Destination User Not Found", "balance": balance.Balance})
		return
	}
	//validate balance
	if balance.Balance < transferRequest.Amount {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Insufficient Balance", "balance": balance.Balance})
		return
	}
	//simulate transfer
	if _, err = service.TopUpBalance(destUser.UserID, transferRequest.Amount); err != nil && err.Error() != "record not found" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Transaction Failed", "balance": balance.Balance})
		return
	}
	//deduction after transfer success
	if _, err = service.DeductBalance(uint64(userID), transferRequest.Amount); err != nil && err.Error() != "record not found" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Transaction Failed", "balance": balance.Balance})
		return
	}
	go service.InsertTransaction(models.Transaction{
		From:   uint64(userID),
		To:     destUser.UserID,
		Amount: transferRequest.Amount,
	})
	c.JSON(http.StatusOK, gin.H{"message": "Transfer Success"})
	return
}
