package database

import "time"

type Task struct {
	ID    int `gorm:"autoincrement"`
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
