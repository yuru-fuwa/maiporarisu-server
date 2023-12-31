package main

import (
	//"app/handler"

	"sqlite/pkg/config"
	"sqlite/pkg/database"
	"sqlite/pkg/handler"
	echovalidator "sqlite/pkg/validator"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.New()
	db, err := database.NewDB(cfg)
	if err != nil {
		panic(err)
	}

	taskHandler := handler.NewTaskHandler(db)
	userHandler := handler.NewUserHandler(db)

	e := echo.New()
	e.Validator = echovalidator.New()
	e.Logger.SetLevel(1)
	task := e.Group("/tasks", testMiddleware)
	task.GET("/:user_id", taskHandler.GetTasks)
	task.POST("", taskHandler.CreateTask)
	//task.GET("/:id", taskHandler.GetTask)
	task.PUT("/:id", taskHandler.UpdateTask)
	task.DELETE("/:id", taskHandler.DeleteTask)

	user := e.Group("/users")
	user.GET("", userHandler.GetUsers)
	user.POST("", userHandler.CreateUser)
	user.GET("/:id", userHandler.GetUser)
	user.PUT("/:id", userHandler.UpdateUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func testMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("test", "test")
		return next(c)
	}
}
