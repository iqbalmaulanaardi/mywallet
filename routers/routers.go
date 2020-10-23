package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mywallet/services"
)

func Configure(app *gin.Engine) {
	app.GET("/", services.Welcome)
	auth := app.Group("/auth")
	{
		auth.POST("/login", services.Login)
		auth.POST("/register", services.Register)
		auth.POST("/logout", services.Logout)
	}
	dashboard := app.Group("/dash")
	{
		dashboard.GET("/balance", services.GetBalance)
	}
	transaction := app.Group("/transaction")
	{
		transaction.GET("", services.GetTransaction)
		transaction.POST("/transfer", services.Transfer)
	}
}
