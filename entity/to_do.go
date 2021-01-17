package entity

import (
		"github.com/jinzhu/gorm"
)

type ToDo struct {
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
}