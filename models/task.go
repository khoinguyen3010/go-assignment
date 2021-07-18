package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID"`
}
