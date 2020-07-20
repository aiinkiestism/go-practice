package models

type Post struct {
	ID      uint   `json:"id" binding:"required`
	Content string `jaon:"content" binding`
	User    User   `json:"-" binindg:"required"`
	UserID  uint   `gorm:"not null" json:"user_id"`
}
