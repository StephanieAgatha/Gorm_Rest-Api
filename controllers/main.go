package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-h8/config"
	"rest-api-h8/entity"
	"time"
)

type OrderRequest struct {
	BuyerEmail   string `json:"buyer_email"`
	BuyerAddress string `json:"buyer_address"`
	ProductId    int    `json:"product_id"`
}

func HandlerGetProduct(ctx *gin.Context) {
	var products []entity.Product
	err := config.DB.Find(&products).Error
	//handling error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func HandlerPostProduct(ctx *gin.Context) {
	//nerima data order
	var orderBody OrderRequest
	err := ctx.ShouldBindJSON(&orderBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	//after product has been created,save to db
	var product entity.Product
	result := config.DB.Where("ID = ?", orderBody.ProductId).First(&product)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Messages": "Product Not Found",
		})
		return
	}

	newOrder := entity.Order{
		BuyerEmail:   orderBody.BuyerEmail,
		BuyerAddress: orderBody.BuyerAddress,
		ProductID:    uint(orderBody.ProductId),
		OrderTime:    time.Now(),
	}
	//handle err when creating new order
	err = config.DB.Create(&newOrder).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Messages": "Error",
		})
	}
	ctx.JSON(http.StatusOK, "Success created order")
}

func HandlerGetOrders(ctx *gin.Context) {
	var orders []entity.Order
	if err := config.DB.Preload("Product").Find(&orders).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	//if success
	ctx.JSON(http.StatusOK, orders)
}
