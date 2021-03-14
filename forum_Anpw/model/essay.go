package model

import "github.com/jinzhu/gorm"

type Essay struct {
	gorm.Model
	Title string `gorm:"type:varchar(20);not null"`
	Essay string `gorm:"size:255;not null"`
	Author string `gorm:"type:varchar(20);not null"`
}
