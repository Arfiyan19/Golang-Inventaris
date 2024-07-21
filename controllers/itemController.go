package controllers

import (
	"inventory-system/database"
	"inventory-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	var items []models.Item
	database.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func CreateItem(c *gin.Context) {
	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&item).Updates(input)
	c.JSON(http.StatusOK, item)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	database.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
