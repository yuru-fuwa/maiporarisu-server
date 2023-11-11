package database

import (
	"time"

	"gorm.io/gorm"
)

// type UUIDBaseModel struct {
// 	// UUID が主キーになり、UUID は Postgres が勝手に生成する
// 	ID        string `gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

type Task struct {
	gorm.Model
	ID string `gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
	//ID    int `gorm:"primaryKey,autoincrement"`
	Time  time.Time
	Name  string `gorm:"size:256"`
	Check bool
}
