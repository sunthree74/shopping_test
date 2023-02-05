package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sunthree74/shopping_test/helper"
	"github.com/sunthree74/shopping_test/interfaces"
	"github.com/sunthree74/shopping_test/model"
	"github.com/sunthree74/shopping_test/structs/request"
	"net/http"
	"strconv"
)

type productCategoryHandler struct {
	usecase     interfaces.ProductCategoryUsecase
	userUSecase interfaces.UserUsecase
}

// HandleSeller is a function to initalize seller handler
func HandleProductCategory(usecase interfaces.ProductCategoryUsecase, userUsecase interfaces.UserUsecase) *productCategoryHandler {
	return &productCategoryHandler{usecase: usecase, userUSecase: userUsecase}
}

func (p *productCategoryHandler) GetList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//_, err := p.userUSecase.FindByJWT(ctx)
		//if err != nil {
		//	helper.LogToFile("error.log", ctx.Request.URL.String(), err.Error())
		//	ctx.Error(err).SetType(gin.ErrorTypePrivate)
		//	ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		//	return
		//}

		category, err := p.usecase.GetList(ctx)
		if err != nil {
			var errNotFound *model.ErrNotFound
			if errors.As(err, &errNotFound) {
				ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": make([]int, 0), "total_rows": 0})
				return
			}

			helper.LogToFile("err.log", ctx.Request.URL.String(), err.Error())
			ctx.Error(err).SetType(gin.ErrorTypePrivate)
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": category})
	}
}

func (p *productCategoryHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//_, err := p.userUSecase.FindByJWT(ctx)
		//if err != nil {
		//	helper.LogToFile("error.log", ctx.Request.URL.String(), err.Error())
		//	ctx.Error(err).SetType(gin.ErrorTypePrivate)
		//	ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		//	return
		//}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": nil})
			return
		}

		category, err := p.usecase.FindById(ctx, uint(id))
		if err != nil {
			var errNotFound *model.ErrNotFound
			if errors.As(err, &errNotFound) {
				ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": make([]int, 0)})
				return
			}

			helper.LogToFile("err.log", ctx.Request.URL.String(), err.Error())
			ctx.Error(err).SetType(gin.ErrorTypePrivate)
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": category})
	}
}

func (p *productCategoryHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var form request.ProductCategoryForm
		if err := ctx.ShouldBindJSON(&form); err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "status": http.StatusBadRequest, "message": "Invalid parameter."})
			return
		}

		if form.Name == " " {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "status": http.StatusBadRequest, "message": "Nama tidak boleh kosong"})
			return
		}

		var category model.ProductCategory
		category.Name = form.Name

		_, err := p.usecase.Create(ctx, &category)
		if err != nil {
			var errNotFound *model.ErrNotFound
			if errors.As(err, &errNotFound) {
				ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": make([]int, 0)})
				return
			}

			helper.LogToFile("err.log", ctx.Request.URL.String(), err.Error())
			ctx.Error(err).SetType(gin.ErrorTypePrivate)
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Berhasil menambahkan kategori"})
	}
}
