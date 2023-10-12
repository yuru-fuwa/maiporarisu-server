package main

import (
	//"app/handler"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	//"github.com/labstack/echo/v4"
)

// var tasks []Tasks

// func allTasks() {
// 	task := Tasks{
// 		Id:    "1",
// 		Time:  "21:00",
// 		Name:  "Task1",
// 		Check: false,
// 	}
// 	tasks = append(tasks, task)
// 	task2 := Tasks{
// 		Id:    "2",
// 		Time:  "22:00",
// 		Name:  "Task2",
// 		Check: false,
// 	}
// 	tasks = append(tasks, task2)
// 	fmt.Println("your tasks are", tasks)
// }

// func handleRoutes() {

// 	taskHandler := handler.NewTaskHandler()

// 	e := echo.New()
// 	task := e.Group("/tasks")
// 	task.GET("", taskHandler.GetTasks)
// 	task.POST("", taskHandler.CreateTask)
// 	task.GET("/:id", taskHandler.GetTask)
// 	task.PUT("/:id", taskHandler.UpdateTask)
// 	task.DELETE("/:id", taskHandler.DeleteTask)

// 	e.Logger.Fatal(e.Start(":8080"))

// 	router := mux.NewRouter()
// 	router.HandleFunc("/gettask/{id}", taskHandler.GetTask).Methods(("GET"))
// 	router.HandleFunc("/create", taskHandler.CreateTask).Methods(("POST"))
// 	router.HandleFunc("/delete/{id}", taskHandler.DeleteTask).Methods(("DELETE"))
// 	router.HandleFunc("/update/{id}", taskHandler.UpdateTask).Methods(("PUT"))
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

func main() {
	db, err := sql.Open("sqlite3", "example.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTaskTableSQL := `
        CREATE TABLE IF NOT EXISTS task (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
			time TEXT,
            name TEXT,
            status INTEGER
        )
    `
	_, err = db.Exec(createTaskTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	insertSQL := "INSERT INTO task (time, name, status) VALUES (?, ?, ?)"
	_, err = db.Exec(insertSQL, "hoge", "hoge", 0)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, time, name, status FROM task")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var time string
		var name string
		var status int
		err = rows.Scan(&id, &time, &name, &status)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Time: %s, Name: %s, Check: %d", id, time, name, status)
	}
	// taskHandler := handler.NewTaskHandler()

	// e := echo.New()
	// task := e.Group("/tasks")
	// task.GET("", taskHandler.GetTasks)
	// task.POST("", taskHandler.CreateTask)
	// task.GET("/:id", taskHandler.GetTask)
	// task.PUT("/:id", taskHandler.UpdateTask)
	// task.DELETE("/:id", taskHandler.DeleteTask)

	// e.Logger.Fatal(e.Start(":8080"))
}
