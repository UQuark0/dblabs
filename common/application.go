package common

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Application struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func NewApplication(dsn string) (*Application, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	router := gin.New()
	router.StaticFile("/lib.js", "./common/lib.js")
	return &Application{
		DB:     db,
		Router: router,
	}, nil
}
