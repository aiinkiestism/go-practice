package models

type User struct {
	ID    uint   `json:"id" binding:"required"`
	Name  string `json:"name" binding:"requied"`
	Posts []Post `json:"posts"`
}
