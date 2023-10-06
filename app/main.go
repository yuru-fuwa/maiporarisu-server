package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Tasks struct {
	Id    string `json:"id"`
	Time  string `json:"time"`
	Name  string `json:"name"`
	Check bool   `json:"check"`
}

var tasks []Tasks

func allTasks() {
	task := Tasks{
		Id:    "1",
		Time:  "21:00",
		Name:  "Task1",
		Check: false,
	}
	tasks = append(tasks, task)
	task2 := Tasks{
		Id:    "2",
		Time:  "22:00",
		Name:  "Task2",
		Check: false,
	}
	tasks = append(tasks, task2)
	fmt.Println("your tasks are", tasks)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am home")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	taskID := mux.Vars(r)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if taskID["id"] == tasks[i].Id {
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}
	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Tasks
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.Id = strconv.Itoa(rand.Intn(1000))
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag := false
	for index, item := range tasks {
		if item.Id == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			flag = true
			json.NewEncoder(w).Encode(map[string]string{"status": "Success"})
			return
		}
	}
	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag := false
	for index, item := range tasks {
		if item.Id == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Tasks
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.Id = params["id"]
			tasks = append(tasks, task)
			flag = true
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods(("GET"))
	router.HandleFunc("/gettasks", getTasks).Methods(("GET"))
	router.HandleFunc("/gettask/{id}", getTask).Methods(("GET"))
	router.HandleFunc("/create", createTask).Methods(("POST"))
	router.HandleFunc("/delete/{id}", deleteTask).Methods(("DELETE"))
	router.HandleFunc("/update/{id}", updateTask).Methods(("PUT"))
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	allTasks()
	fmt.Println("Hello World")
	handleRoutes()
}
