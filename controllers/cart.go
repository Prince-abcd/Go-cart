package controllers

import (
	"Cart/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Cartcontrollers struct{}

func (ct *Cartcontrollers) Cartmain(r *gin.Engine) {
	cart := r.Group("/cart")
	order:=r.Group("/odrer")
	cart.POST("", ct.AddItem)
	cart.GET("", ct.GetItem)
	order.POST("")
}

// adding the cart item into user cart if user doesnt have cart we will create if already have we will update that item arary in the cart
func (ct *Cartcontrollers) AddItem(c *gin.Context) {
	var requestBody struct {
		UserID uint          `json:"user_id"`
		Item   database.Item `json:"item"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var cart database.Cart
	if err := database.DB.Preload("Items").Where("user_id = ?", requestBody.UserID).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			cart = database.Cart{UserID: requestBody.UserID}
			if err := database.DB.Create(&cart).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find cart"})
			return
		}
	}

	var existingItem database.Item
	if err := database.DB.Where("name = ?", requestBody.Item.Name).First(&existingItem).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			if err := database.DB.Create(&requestBody.Item).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
				return
			}
			existingItem = requestBody.Item
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find item"})
			return
		}
	} else {

		existingItem = requestBody.Item
	}

	cart.Items = append(cart.Items, existingItem)
	if err := database.DB.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart", "cart": cart})
}

//getting all the cart items of the user based on the user id 

func (ct *Cartcontrollers) GetItem(c *gin.Context) {
	var requestBody struct {
		UserID uint `json:"user_id"`
	}


	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	var cart database.Cart
	if err := database.DB.Preload("Items").Where("user_id = ?", requestBody.UserID).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found for user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{"items": cart.Items})
}

//crating order from cart
func (ct *Cartcontrollers) CreateOrder(c *gin.Context) {
	var requestBody struct {
		UserID uint           `json:"user_id"`
		Items  []database.Item `json:"items"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user database.User
	if err := database.DB.First(&user, requestBody.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	order := database.Order{
		UserID: requestBody.UserID,
		Items:  requestBody.Items,
	}

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order": order})
}