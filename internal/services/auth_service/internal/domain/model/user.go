package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `validate:"uuid" gorm:"not null"`
	Email    string `validate:"email" gorm:"not null"`
	Role     string `validate:"uuid" gorm:"not null"`
	Password string `gorm:"not null"`
}
