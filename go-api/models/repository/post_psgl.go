package repository

import (
	"fmt"

	"go/src/go-api/db"
	"go/src/go-api/form/api"
	"go/src/go-api/models"

	"github.com/gin-gonic/gin"
)

type PostRepository struct{}

type Post api.Post

func (_ PostRepository) GetAll() ([]Post, error) {
	db := db.GetDB()
	var p []Post

	if err := db.Find(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (_ PostRepository) CreateModel(p *models.Post) (*models.Post, error) {
	db := db.GetDB()
	if err := db.Create(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (_ PostRepository) GetByID(id int) (models.Post, error) {
	db := db.GetDB()
	var p models.Post
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (_ PostRepository) UpdateByID(id int, c *gin.Context) (api.Post, error) {
	db := db.GetDB()
	var p api.Post
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}
	userID := p.UserID
	if err := c.BindJSON(&p); err != nil {
		return p, err
	}
	fmt.Printf("%+V", p)
	p.ID = uint(id)
	p.UserID = userID
	db.Save(&p)

	return p, nil
}

func (_ PostRepository) DeleteByID(id int) error {
	db := db.GetDB()
	var p Post

	if err := db.Where("id = ?", id).Delete(&p).Error; err != nil {
		return err
	}

	return nil
}
