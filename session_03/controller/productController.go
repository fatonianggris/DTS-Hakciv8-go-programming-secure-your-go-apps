package controller

import (
	"go-programming-secure-your-go-apps/session_03/database"
	"go-programming-secure-your-go-apps/session_03/entity"
	"go-programming-secure-your-go-apps/session_03/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProductsOwnedByUserId(ctx *gin.Context) {
	db := database.GetDB()
	userId := helper.GetUserIdFromClaims(ctx)
	status := http.StatusFound
	products := []entity.Product{}

	errFind := db.Where("user_id = ?", userId).Find(&products).Error
	if errFind != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseError(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessWithData(status, products))
}

func GetProductById(ctx *gin.Context) {
	db := database.GetDB()
	status := http.StatusFound
	product := entity.Product{}

	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseError(status))
		return
	}

	errFind := db.First(&product, "id = ?", id).Error
	if errFind != nil {
		status = http.StatusNotFound
		ctx.AbortWithStatusJSON(status, helper.ResponseError(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessWithData(status, product))
}

func CreateNewProduct(ctx *gin.Context) {
	db := database.GetDB()
	userId := helper.GetUserIdFromClaims(ctx)
	status := http.StatusCreated
	newProduct := entity.Product{}

	errBind := ctx.ShouldBindJSON(&newProduct)
	if errBind != nil {
		status := http.StatusBadRequest
		ctx.AbortWithStatusJSON(status, helper.ResponseError(status))
		return
	}

	newProduct.UserId = userId

	errCreate := db.Create(&newProduct).Error
	if errCreate != nil {
		status = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(status, helper.ResponseError(status))
		return
	}

	ctx.JSON(status, helper.ResponseSuccessNoData(status))
}
