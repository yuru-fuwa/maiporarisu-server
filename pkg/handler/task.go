package handler

import (
	"log"
	"net/http"
	"sqlite/pkg/database"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	Task struct {
		gorm.Model
		ID    string `gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
		Time  time.Time
		Name  string `gorm:"size:256"`
		Check bool
	}

	GetTasksRequest  struct{}
	GetTasksResponse struct {
		Tasks []Task `json:"tasks"`
	}

	GetTaskRequest struct {
		ID string `json:"id"`
	}
	GetTaskResponse struct {
		ID    string `json:"id"`
		Time  string `json:"time"`
		Name  string `json:"Name"`
		Check string `json:"check"`
	}

	CreateTaskRequest struct {
		Time string `json:"time"`
		Name string `json:"name"`
	}
	CreateTaskResponse struct{}

	DeleteTaskRequest struct {
		ID string `param:"id"`
	}
	DeleteTaskResponse struct{}

	UpdateTaskRequest struct {
		ID    string `param:"id"`
		Time  string `json:"time"`
		Name  string `json:"name"`
		Check bool   `json:"check"`
	}
	UpdateTaskResponse struct {
		Time  string `json:"time"`
		Name  string `json:"name"`
		Check bool   `json:"check"`
	}
)

type taskHandler struct {
	db *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *taskHandler {
	return &taskHandler{
		db: db,
	}
}

func (h *taskHandler) GetTasks(c echo.Context) error {
	// tasks := map[string]string{"hoge": "fuga"}
	// return c.JSON(http.StatusOK, tasks)
	tasks := []Task{}

	if err := h.db.Find(&tasks).Error; err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, "failed to get tasks")
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *taskHandler) GetTask(c echo.Context) error {
	// tasks := map[string]string{"hoge": "fuga"}
	// taskID := mux.Vars(r)
	// flag := false
	// for i := 0; i < len(tasks); i++ {
	// 	if taskID["id"] == tasks[i].ID {
	// 		json.NewEncoder(w).Encode(tasks[i])
	// 		flag = true
	// 		break
	// 	}
	// }
	// if !flag {
	// 	json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	// }
	return nil
}

func (h *taskHandler) CreateTask(c echo.Context) error {
	task := &CreateTaskRequest{}
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	t, err := time.Parse("2006-01-02 15:04:05", task.Time)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	result := h.db.Create(&database.Task{
		Name:  task.Name,
		Time:  t,
		Check: false,
	})

	return result.Error
}

func (h *taskHandler) DeleteTask(c echo.Context) error {
	task := &DeleteTaskRequest{}
	if err := c.Bind(task); err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	log.Print("TaskID:" + task.ID)

	deTask := &database.Task{}
	if err := h.db.Where("id = ?", task.ID).Delete(&deTask).Error; err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, "failed to delete task")
	}
	log.Print("delete task")
	return c.JSON(http.StatusOK, "success")
}

// w.Header().Set("Content-Type", "application/json")
// params := mux.Vars(r)
// flag := false
// for index, item := range tasks {
// 	if item.ID == params["id"] {
// 		tasks = append(tasks[:index], tasks[index+1:]...)
// 		flag = true
// 		json.NewEncoder(w).Encode(map[string]string{"status": "Success"})
// 		return
// 	}
// }
// if !flag {
// 	json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
// }

func (h *taskHandler) UpdateTask(c echo.Context) error {
	task := &UpdateTaskRequest{}
	if err := c.Bind(task); err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	log.Print("TaskID:" + task.ID)

	t, err := time.Parse("2006-01-02 15:04:05", task.Time)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	upTask := &database.Task{}
	upTask.ID = task.ID
	upTask.Name = task.Name
	upTask.Time = t
	upTask.Check = task.Check

	if err := h.db.Where("id = ?", task.ID).Save(&upTask).Error; err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, "failed to update task")
	}

	log.Print("update task")
	return c.JSON(http.StatusOK, "success")
}

// w.Header().Set("Content-Type", "application/json")
// params := mux.Vars(r)
// flag := false
// for index, item := range tasks {
// 	if item.ID == params["id"] {
// 		tasks = append(tasks[:index], tasks[index+1:]...)
// 		var task Tasks
// 		_ = json.NewDecoder(r.Body).Decode(&task)
// 		task.ID = params["id"]
// 		tasks = append(tasks, task)
// 		flag = true
// 		json.NewEncoder(w).Encode(task)
// 		return
// 	}
// }
// if !flag {
// 	json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
// }
