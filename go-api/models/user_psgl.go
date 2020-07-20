package repotitory

import (
	"github.com/gin-gonic/gin"
	"go/src/go-api/db"
	"go/src/go-api/models"
)

type UserRepository struct{}

type User models.User

type UserProfile struct {
	Name string
	Id int
}

func (_ UserRepository) GetAll() ([]UserProfile, error) {
	db := db.GetDB()
	var u []UserProfile
	if err := db.Table("users").Select("name, id").Scan(&u).Error; err != nil {
		return nil err
	}
	return u, nil
}

func (_ UserRepository) CreateModel(c *gin.Context) (user, error) {
	db := db.GetDB()
	var u User
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}
	if err := db.Create(&u).Error; er != nil {
		return u, err
	}
	return u, nil
}

func (_ UserRepository) GetByID(id int) (models.User, error) {
	db := db.GetDB()
	var me models.User
	if err := db.Where("id = ?", id).First(&me).Error; err != nil {
		return me, err
	}
	var posts []models.Post
	db.Where("id = ?", id).First(&me)
	db.Model(&me).Related(&posts)
	me.Posts = posts

	return me, nil
}

func (_ UserRepository) UpdateByID(ID, int, c *gin.Context) (models.User, error) {
	db := db.GetDB()
	var u models.User
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}
	u.ID = uint(id)
	db.Save(&u)

	return u, nil
}

func (_ UserRepository) DeleteByID(id int) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}