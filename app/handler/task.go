package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Task struct {
		Id    string `json:"id"`
		Time  string `json:"time"`
		Name  string `json:"name"`
		Check bool   `json:"check"`
	}

	GetTasksRequest  struct{}
	GetTasksResponse struct {
		Tasks []Task `json:"tasks"`
	}

	GetTaskRequest struct {
		Id string `json:"id"`
	}
	GetTaskResponse struct {
		Id string `json:"id"`
		Time string `json:"time"`
		Name string `json:"Name"`
		Check string `json:"check"`
	}

	CreateTaskRequest struct {
		Time string `json:"time"`
		Name string `json:"name"`
	}
	CreateTaskResponse struct{}

	DeleteTaskRequest struct {
		Id string `param:"id"`
	}
	DeleteTaskResponse struct{}

	UpdateTaskRequest struct {
		Id   string `param:"id"`
		Time string `json:"time"`
		Name string `json:"name"`
		Check bool `json:"check"`
	}
	UpdateTaskResponse struct {
		Time  string `json:"time"`
		Name  string `json:"name"`
		Check bool   `json:"check"`
	}
)

type taskHandler struct {
}

func NewTaskHandler() *taskHandler {
	return &taskHandler{}
}

func (h *taskHandler) GetTasks(c echo.Context) error {
	tasks := map[string]string{"hoge": "fuga"}
	return c.JSON(http.StatusOK, tasks)
}

func (h *taskHandler) GetTask(c echo.Context) error {
	// tasks := map[string]string{"hoge": "fuga"}
	// taskID := mux.Vars(r)
	// flag := false
	// for i := 0; i < len(tasks); i++ {
	// 	if taskID["id"] == tasks[i].Id {
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
	task := &Task{}
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	// w.Header().Set("Content-Type", "application/json")
	// var task Tasks
	// _ = json.NewDecoder(r.Body).Decode(&task)
	// task.Id = strconv.Itoa(rand.Intn(1000))
	// tasks = append(tasks, task)
	// json.NewEncoder(w).Encode(task)
	return nil
}

func (h *taskHandler) DeleteTask(c echo.Context) error {
	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// flag := false
	// for index, item := range tasks {
	// 	if item.Id == params["id"] {
	// 		tasks = append(tasks[:index], tasks[index+1:]...)
	// 		flag = true
	// 		json.NewEncoder(w).Encode(map[string]string{"status": "Success"})
	// 		return
	// 	}
	// }
	// if !flag {
	// 	json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	// }
	return nil
}

func (h *taskHandler) UpdateTask(c echo.Context) error {
	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// flag := false
	// for index, item := range tasks {
	// 	if item.Id == params["id"] {
	// 		tasks = append(tasks[:index], tasks[index+1:]...)
	// 		var task Tasks
	// 		_ = json.NewDecoder(r.Body).Decode(&task)
	// 		task.Id = params["id"]
	// 		tasks = append(tasks, task)
	// 		flag = true
	// 		json.NewEncoder(w).Encode(task)
	// 		return
	// 	}
	// }
	// if !flag {
	// 	json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	// }
	return nil
}
