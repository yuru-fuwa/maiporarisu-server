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

// createTaskTableSQL := `
//     CREATE TABLE IF NOT EXISTS task (
//         id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		time TEXT,
//         name TEXT,
//         status INTEGER
//     )
// `
// _, err = db.Exec(createTaskTableSQL)
// if err != nil {
// 	log.Fatal(err)
// }
// insertSQL := "INSERT INTO task (time, name, status) VALUES (?, ?, ?)"
// _, err = db.Exec(insertSQL, "hoge", "hoge", 0)
// if err != nil {
// 	log.Fatal(err)
// }

// rows, err := db.Query("SELECT id, time, name, status FROM task")
// if err != nil {
// 	log.Fatal(err)
// }
// defer rows.Close()

// for rows.Next() {
// 	var id int
// 	var time string
// 	var name string
// 	var status int
// 	err = rows.Scan(&id, &time, &name, &status)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("ID: %d, Time: %s, Name: %s, Check: %d", id, time, name, status)
// }
