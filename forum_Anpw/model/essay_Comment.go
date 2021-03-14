package model

import "github.com/jinzhu/gorm"

type Essay_Comment struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null"`
	Comment string `gorm:"size:255;not null"`
	EssayID string `gorm:"type:varchar(3);not null"`
}
