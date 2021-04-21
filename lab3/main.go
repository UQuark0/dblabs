package main

import (
	"dblabs/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

func Register(app *common.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := common.User{}
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.String(http.StatusBadRequest, "%s", err.Error())
			return
		}
		user := request
		user.GeneratePassword(request.Password)
		err = app.DB.Create(&user).Error
		if err != nil {
			ctx.String(http.StatusInternalServerError, "%s", err.Error())
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func Login(app *common.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := common.User{}
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.String(http.StatusBadRequest, "%s", err.Error())
			return
		}
		user := common.User{}
		err = app.DB.Where("username = ?", request.Username).First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.String(http.StatusNotFound, "%s", err.Error())
			} else {
				ctx.String(http.StatusInternalServerError, "%s", err.Error())
			}
			return
		}
		if !user.ValidatePassword(request.Password) {
			ctx.String(http.StatusUnauthorized, "invalid password")
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func GetUsers(app *common.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var users []*common.User
		err := app.DB.Find(&users).Error
		if err != nil {
			ctx.String(http.StatusInternalServerError, "%s", err.Error())
			return
		}
		ctx.JSON(http.StatusOK, users)
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	app, err := common.NewApplication("labber:labber@tcp(localhost:3306)/dblabs?parseTime=true")
	if err != nil {
		panic(err)
	}

	app.Router.StaticFile("/", "./lab3/index.html")
	app.Router.POST("/register", Register(app))
	app.Router.POST("/login", Login(app))
	app.Router.GET("/users", GetUsers(app))

	err = app.Router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

