package services

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mywallet/models"
	"github.com/iqbalmaulanaardi/mywallet/repository"
	"github.com/iqbalmaulanaardi/mywallet/responses"
)

func GetBalance(c *gin.Context) {
	var (
		err                 error
		claims              responses.MyClaims
		ok                  bool
		authorizationHeader string
		balance             models.Balance
	)
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	authorizationHeader = c.Request.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	if token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return []byte("9labqi6"), nil
	}); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	claims2, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(claims2)
	ctx := context.WithValue(context.Background(), "user_info", claims2)
	userInfo := ctx.Value("user_info").(jwt.MapClaims)
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
