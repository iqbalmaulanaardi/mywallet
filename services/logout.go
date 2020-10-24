package services

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mywallet/repository"
)

func Logout(c *gin.Context) {
	var (
		authorizationHeader string
		userInfo            map[string]interface{}
		err                 error
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
	fmt.Println(userID)
	if _, err = service.Logout(uint64(userID)); err != nil && err.Error() != "record not found" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logout success"})
	return
}
