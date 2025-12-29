package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	var u User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request Body",
		})

		return
	}

	if err := CreateUser(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    u,
	})
}

func GetUserHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "User service ok",
	})
}
