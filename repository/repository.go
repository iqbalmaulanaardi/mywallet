package repository

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/iqbalmaulanaardi/mywallet/models"
	"github.com/jinzhu/gorm"
	"go.elastic.co/apm/module/apmgorm"
	_ "go.elastic.co/apm/module/apmgorm/dialects/postgres"
)

type Service struct {
	DB *gorm.DB
}

func Configure() (service Service, err error) {
	DbConfig := "host=" + "127.0.0.1" + " port=" + "5432" + " user=" + "postgres" + " dbname=" + "mywallet" + " password=" + "postgres" + " sslmode=disable"
	if service.DB, err = apmgorm.Open("postgres", DbConfig); err != nil {
		return
	}
	ctx := context.Background()
	service.DB = apmgorm.WithContext(ctx, service.DB)
	sqlDB := service.DB.DB()
	sqlDB.SetMaxIdleConns(5)

	return
}
func GinHandler(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("service", s)
		c.Next()
	}
}
func (s *Service) AutoMigrate() {
	s.DB.AutoMigrate(models.User{}, models.Balance{}, models.Session{}, models.Transaction{})
	s.DB.Model(&models.Balance{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
	s.DB.Model(&models.Session{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
	s.DB.Model(&models.Transaction{}).AddForeignKey("from", "users(user_id)", "RESTRICT", "RESTRICT")
	s.DB.Model(&models.Transaction{}).AddForeignKey("to", "users(user_id)", "RESTRICT", "RESTRICT")

}
