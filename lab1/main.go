package main

import (
	"dblabs/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(app *common.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := struct {
			Products []*common.Product
			Fats     []*common.Fat
			Infos    []*common.Info
			Sales    []*common.Sale
		}{
			Products: make([]*common.Product, 0),
			Fats:     make([]*common.Fat, 0),
			Infos:    make([]*common.Info, 0),
			Sales:    make([]*common.Sale, 0),
		}

		err := app.DB.Find(&response.Products).Error
		if err != nil {
			ctx.String(http.StatusInternalServerError, "%s", err.Error())
			return
		}
		err = app.DB.Find(&response.Fats).Error
		if err != nil {
			ctx.String(http.StatusInternalServerError, "%s", err.Error())
			return
		}
		err = app.DB.Find(&response.Infos).Error
		if err != nil {
			ctx.String(http.StatusInternalServerError, "%s", err.Error())
			return
		}
		err = app.DB.Find(&response.Sales).Error
		if err != nil {
			ctx.String(http.StatusInternalServerError, "%s", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}

func main() {
	app, err := common.NewApplication("labber:labber@tcp(localhost:3306)/dblabs?parseTime=true")
	if err != nil {
		panic(err)
	}

	app.Router.StaticFile("/", "./lab1/index.html")
	app.Router.GET("/all", GetAll(app))

	err = app.Router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
