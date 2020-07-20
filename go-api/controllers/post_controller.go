package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go/src/go-api/form/api"
	"go/src/go-api/models"
	"go/src/go-api/models/repository"
)

type PostController struct{}

func (pc PostController) Index(c *gin.Context) {
	var u repository.PostRepository
	p, err := u.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

func (pc PostController) Create(c *gin.Context) {
	var u repository.PostRepository
	in := api.Post{}
	if err := c.BindJSON(&in); err != nil {
		return
	}
	in2 := &models.Post{
		ID: in.ID,
		Content: in.Content,
		UserID: in.UserID,
	}
	p, err := u.CreateModel(in2)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, p)
	}
}

func (pc PostController) Show (c &gin.Context) {
	id := c.Params.ByName("id")
	var p repository.PostRepository
	idInt, _ := strconv.Atoi(id)
	post, err := p.GetByID(idInt)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, post)
	}
}

func (pc PostController) Update(c *gin.Context) {
	id := c.Parmas.ByName("id")
	var u repository.PostRepository
	idInt, _ := strconv.Atoi(id)
	p, err := u.UpdateById(idInt, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

func (pc PostController) Delete(c *gin.Context) {
	id := c.Parmas.ByName("id")
	var u repository.PostRepository
	idInt, _ := strconv.Atoi(id)
	if err := u.DeleteById(idInt); err != nil {
		c.AbortWithStatus(403)
		c.JSON(htp.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "ID" + id + "の投稿を削除しました"})
}
