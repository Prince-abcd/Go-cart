package controllers

import (
	"net/http"

	"Cart/database"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Basecontrollers struct{}

func (b *Basecontrollers) Dostuff(r *gin.Engine) {
	user := r.Group("/users")
	user.POST("", b.Adduser)
	user.GET("", b.Getuser)
	user.POST("/login", b.Getloginuser)
}

func (b *Basecontrollers) Adduser(c *gin.Context) {
	var user = database.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = string(hashedPassword)
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
func (b *Basecontrollers) Getuser(c *gin.Context) {
	var users []database.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
func (b *Basecontrollers) Getloginuser(c *gin.Context) {
	var userInfo = database.User{}
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user = database.User{}
	if err := database.DB.Where("username = ?", userInfo.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is invalid"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInfo.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": user.Username, "data": user})
}
