package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string `validate:"uuid" gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `validate:"email" gorm:"not null"`
	Password  string `gorm:"not null"`
	Role      string `validate:"uuid" gorm:"not null"`
}
