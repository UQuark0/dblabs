package main

import (
	"dblabs/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetModel(_ *common.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := struct {
			Name string `form:"name" binding:"required"`
		}{}

		err := ctx.ShouldBindQuery(&request)
		if err != nil {
			ctx.String(http.StatusBadRequest, "%s", err.Error())
			return
		}

		var model interface{}

		switch request.Name {
		case "fat":
			model = common.Fat{}
		case "info":
			model = common.Info{}
		case "product":
			model = common.Product{}
		case "sale":
			model = common.Sale{}
		default:
			ctx.String(http.StatusNotFound, "%s", "model not found")
			return
		}

		ctx.JSON(http.StatusOK, model)
	}
}

func Select(app *common.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := struct {
			Model string `uri:"model" binding:"required"`
		}{}

		err := ctx.ShouldBindUri(&uri)
		if err != nil {
			ctx.String(http.StatusBadRequest, "%s", err.Error())
			return
		}

		var model interface{}

		switch uri.Model {
		case "fat":
			fat := common.Fat{}
			err = ctx.ShouldBindJSON(&fat)
			if err != nil {
				ctx.String(http.StatusBadRequest, "%s", err.Error())
				return
			}
			err = app.DB.First(&fat).Error
			model = fat
		case "info":
			info := common.Info{}
			err = ctx.ShouldBindJSON(&info)
			if err != nil {
				ctx.String(http.StatusBadRequest, "%s", err.Error())
				return
			}
			err = app.DB.First(&info).Error
			model = info
		case "product":
			product := common.Product{}
			err = ctx.ShouldBindJSON(&product)
			if err != nil {
				ctx.String(http.StatusBadRequest, "%s", err.Error())
				return
			}
			err = app.DB.First(&product).Error
			model = product
		case "sale":
			sale := common.Sale{}
			err = ctx.ShouldBindJSON(&sale)
			if err != nil {
				ctx.String(http.StatusBadRequest, "%s", err.Error())
				return
			}
			err = app.DB.First(&sale).Error
			model = sale
		default:
			ctx.String(http.StatusNotFound, "%s", "model not found")
			return
		}

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.String(http.StatusNotFound, "%s", err.Error())
			} else {
				ctx.String(http.StatusInternalServerError, "%s", err.Error())
			}
			return
		}

		ctx.JSON(http.StatusOK, model)
	}
}

func Insert(app *common.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := struct {
			Model string `uri:"model" binding:"required"`
		}{}

		err := ctx.ShouldBindUri(&uri)
		if err != nil {
			ctx.String(http.StatusBadRequest, "%s", err.Error())
			return
		}

		var model interface{}

		switch uri.Model {
		case "fat":
			fat := common.Fat{}
			err = ctx.ShouldBindJSON(&fat)
			if err != nil {
				ctx.String(http.StatusBadRequest, "%s", err.Error())
				return
			}
			err = app.DB.Create(&fat).Error
			model = fat
		case "info":
			info := common.Info{}
			err = ctx.ShouldBindJSON(&info)
			if err != nil {
				ctx.String(http.StatusBadRequest, "%s", err.Error())
				return
			}
			err = app.DB.Create(&info).Error
			model = info
		case "product":
			product := common.Product{}
			err = ctx.ShouldBindJSON(&product)
			if err != nil {
				ctx.String(http.StatusBadRequest, "%s", err.Error())
				return
			}
			err = app.DB.Create(&product).Error
			model = product
		case "sale":
			sale := common.Sale{}
			err = ctx.ShouldBindJSON(&sale)
			if err != nil {
				ctx.String(http.StatusBadRequest, "%s", err.Error())
				return
			}
			err = app.DB.Create(&sale).Error
			model = sale
		default:
			ctx.String(http.StatusNotFound, "%s", "model not found")
			return
		}

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.String(http.StatusNotFound, "%s", err.Error())
			} else {
				ctx.String(http.StatusInternalServerError, "%s", err.Error())
			}
			return
		}

		ctx.JSON(http.StatusOK, model)
	}
}

func Update(app *common.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := struct {
			Model string `uri:"model" binding:"required"`
		}{}

		err := ctx.ShouldBindUri(&uri)
		if err != nil {
			ctx.String(http.StatusBadRequest, "%s", err.Error())
			return
		}

		var model interface{}

		switch uri.Model {
		case "fat":
			fat := common.Fat{}
			err = ctx.ShouldBindJSON(&fat)
			if err != nil {
				ctx.String(http.StatusBadRequest, "%s", err.Error())
				return
			}
			err = app.DB.Save(&fat).Error
			model = fat
		case "info":
			info := common.Info{}
			err = ctx.ShouldBindJSON(&info)
			if err != nil {
				ctx.String(http.StatusBadRequest, "%s", err.Error())
				return
			}
			err = app.DB.Save(&info).Error
			model = info
		case "product":
			product := common.Product{}
			err = ctx.ShouldBindJSON(&product)
			if err != nil {
				ctx.String(http.StatusBadRequest, "%s", err.Error())
				return
			}
			err = app.DB.Save(&product).Error
			model = product
		case "sale":
			sale := common.Sale{}
			err = ctx.ShouldBindJSON(&sale)
			if err != nil {
				ctx.String(http.StatusBadRequest, "%s", err.Error())
				return
			}
			err = app.DB.Save(&sale).Error
			model = sale
		default:
			ctx.String(http.StatusNotFound, "%s", "model not found")
			return
		}

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.String(http.StatusNotFound, "%s", err.Error())
			} else {
				ctx.String(http.StatusInternalServerError, "%s", err.Error())
			}
			return
		}

		ctx.JSON(http.StatusOK, model)
	}
}

func main() {
	app, err := common.NewApplication("labber:labber@tcp(localhost:3306)/dblabs?parseTime=true")
	if err != nil {
		panic(err)
	}

	app.Router.StaticFile("/", "./lab2/index.html")
	app.Router.GET("/model", GetModel(app))
	app.Router.POST("/:model/select", Select(app))
	app.Router.POST("/:model/insert", Insert(app))
	app.Router.POST("/:model/update", Update(app))

	err = app.Router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
