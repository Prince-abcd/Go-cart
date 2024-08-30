package controllers

import (
	"Cart/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Itemcontrollers struct{}

func (i *Itemcontrollers) Itemmain(r *gin.Engine) {
	Item := r.Group("/item")
	Item.POST("", i.AddItem)
	Item.GET("", i.GetItem)
}

func (i *Itemcontrollers) AddItem(c *gin.Context) {
	var item database.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item created", "item": item})
}
func (i *Itemcontrollers) GetItem(c *gin.Context) {
	var items []database.Item
	if err := database.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve items"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}
