package main

import (
	//"app/handler"

	"sqlite/pkg/config"
	"sqlite/pkg/database"
	"sqlite/pkg/handler"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

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
