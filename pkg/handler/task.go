package handler

import (
	"errors"
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

	SimplifiedTask struct {
		ID    string    `json:"id"`
		Time  time.Time `json:"time"`
		Name  string    `json:"name"`
		Check bool      `json:"check"`
	}

	GetTasksRequest  struct{}
	GetTasksResponse struct {
		Tasks []SimplifiedTask `json:"tasks"`
	}

	GetTaskRequest struct {
		ID string `param:"id"`
	}
	GetTaskResponse struct {
		ID    string `json:"id"`
		Time  string `json:"time"`
		Name  string `json:"Name"`
		Check string `json:"check"`
	}

	CreateTaskRequest struct {
		Time string `json:"time" validator:"required,datetime=2006-01-02T15:04:05Z07:00"`
		Name string `json:"name"`
	}
	CreateTaskResponse struct{}

	DeleteTaskRequest struct {
		ID string `param:"id"`
	}
	DeleteTaskResponse struct{}

	UpdateTaskRequest struct {
		ID    string `param:"id" validator:"required,uuid"`
		Time  string `json:"time" validator:"required,datetime=2006-01-02T15:04:05Z07:00"`
		Name  string `json:"name"`
		Check bool   `json:"check"`
	}
	UpdateTaskResponse struct {
		Time  string `json:"time" validator:"required,datetime=2006-01-02T15:04:05Z07:00"`
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
	simplifiedTasks := []SimplifiedTask{}

	if err := h.db.Table("tasks").Select("id", "time", "name", "check").Find(&simplifiedTasks).Error; err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "failed to get tasks")
	}
	return c.JSON(http.StatusOK, simplifiedTasks)
}

func (h *taskHandler) GetTask(c echo.Context) error {
	req := &GetTaskRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	task := &Task{}
	if err := h.db.Where("id = ?", req.ID).First(&task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "task not found")
		}
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, "failed to get task")
	}

	return c.JSON(http.StatusOK, task)
}

func (h *taskHandler) CreateTask(c echo.Context) error {
	task := &CreateTaskRequest{}
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	if err := c.Validate(task); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	t, err := time.Parse(time.RFC3339, task.Time)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid time format")
	}

	dbTask := &database.Task{
		Name:  task.Name,
		Time:  t,
		Check: false,
	}

	if err := h.db.Create(dbTask).Error; err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "failed to create task")
	}

	return c.JSON(http.StatusOK, dbTask)
}

func (h *taskHandler) DeleteTask(c echo.Context) error {
	task := &DeleteTaskRequest{}
	if err := c.Bind(task); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	log.Print("TaskID:" + task.ID)

	if err := h.db.Where("id = ?", task.ID).Delete(&database.Task{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "task not found")
		}
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, "failed to delete task")
	}
	log.Print("delete task")
	return c.JSON(http.StatusOK, "success")
}

func (h *taskHandler) UpdateTask(c echo.Context) error {
	task := &UpdateTaskRequest{}
	if err := c.Bind(task); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	log.Print("TaskID:" + task.ID)

	if err := c.Validate(task); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}
	t, err := time.Parse(time.RFC3339, task.Time)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid time format")
	}

	upTask := &database.Task{
		ID:    task.ID,
		Name:  task.Name,
		Time:  t,
		Check: task.Check,
	}

	if err := h.db.Where("id = ?", task.ID).Save(&upTask).Error; err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "failed to update task")
	}

	log.Print("update task")
	return c.JSON(http.StatusOK, "success")
}
