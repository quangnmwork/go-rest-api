package main

import (
	"fmt"
	"net/http"
	"todo-app/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Run migrations
	if err := utils.AutoMigrate(); err != nil {
		fmt.Printf("Migration error: %v\n", err)
		panic("Failed to migrate database")
	}

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, World! 3",
		})
	})

	router.Run(":4000")
}
