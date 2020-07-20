package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go/src/go-api/models/repository"
)

type UserController struct{}

func (pc UserController) Index(c *gin.Context) {
	var u repository.UserRepository
	p, err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

func (pc UserController) Create(c *gin.Context) {
	var u repository.UserRepository
	p, err := u.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, p)
	}
}

func (pc UserController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.UserRepository
	idInt, _ := strconv.Atoi(id)
	p, err := u.UpdateByID(idInt, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

func (pc UserController) Delete(c *gin.Context) {
	id := c.Params.ByName("Id")
	var u repository.UserRepository
	idInt, _ := strconv.Atoi(id)
	if err := u.DeelteByID(idInt); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "ID" + id + "のユーザーを削除しました"})
}