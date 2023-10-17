package main

import (
	//"app/handler"

	"sqlite/pkg/config"
	"sqlite/pkg/database"
	"sqlite/pkg/handler"

	"github.com/labstack/echo/v4"
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
	cfg := config.New()
	db, err := database.NewDB(cfg)
	if err != nil {
		panic(err)
	}

	taskHandler := handler.NewTaskHandler(db)

	e := echo.New()
	task := e.Group("/tasks")
	task.GET("", taskHandler.GetTasks)
	task.POST("", taskHandler.CreateTask)
	task.GET("/:id", taskHandler.GetTask)
	task.PUT("/:id", taskHandler.UpdateTask)
	task.DELETE("/:id", taskHandler.DeleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}
