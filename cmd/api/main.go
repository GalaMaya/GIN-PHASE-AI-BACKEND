package main

import (
	"github.com/GalaMaya/GIN-PHASE-AI-BACKEND.git/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.POST("/users", user.CreateUserHandler)

	r.GET("/users/health", user.GetUserHealthHandler)

	r.Run(":8080") // listen and serve on 8080
}
