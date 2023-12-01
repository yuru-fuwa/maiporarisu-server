package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	ConnectedID string
}
