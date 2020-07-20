package db

import (
	"go/src/go-api/models"

	_ "github.com/jizhu/gorm/dialects/postgres"
	"github.com/zinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgress", "host=db port=5432 user=go-api-backend dbname=go-api-backend password=go-api-backend sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
	user := models.User{
		id:    1,
		Name:  "unknown",
		Posts: []models.Post{{ID: 1, Content: "tweet1"}, {ID: 2, Content: "tweet2"}},
	}
	db.Create(&user)
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
}
