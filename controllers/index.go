package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kubil-ismail/go-example/models"
	"gorm.io/gorm"
)

func Index(connection *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	connection.JSON(http.StatusOK, gin.H{"products": products})
}

func Detail(connection *gin.Context) {
	var productId = connection.Param("id")
	var products models.Product

	if err := models.DB.First(&products, productId).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			connection.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return
		default:
			connection.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	connection.JSON(http.StatusOK, gin.H{"Products": products})
}

func Add(connection *gin.Context) {
	var products models.Product

	if err := connection.ShouldBindJSON(&products); err != nil {
		connection.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&products)

	connection.JSON(http.StatusOK, gin.H{"Products": products})
}

func Update(connection *gin.Context) {
	var productId = connection.Param("id")
	var products models.Product

	if err := connection.ShouldBindJSON(&products); err != nil {
		connection.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&products).Where("id = ?", productId).Updates(&products).RowsAffected == 0 {
		connection.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Cannot update product"})
		return
	}

	connection.JSON(http.StatusOK, gin.H{"Products": products, "message": "Data success updated"})
}

func Delete(connection *gin.Context) {
	var products models.Product
	var productId = connection.Param("id")

	if models.DB.Delete(&products, productId).RowsAffected == 0 {
		connection.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Error on server"})
		return
	}

	connection.JSON(http.StatusOK, gin.H{"message": "Data success updated"})
}
