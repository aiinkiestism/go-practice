package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_test2/controllers"
	"gin_test2/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
