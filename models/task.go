package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Content string `json:"content"`
	UserID  User   `json:"user_id" gorm:"foreignKey:id"`
}
