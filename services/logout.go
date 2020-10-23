package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
	return
}
