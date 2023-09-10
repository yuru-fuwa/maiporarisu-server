package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Tasks struct {
	TaskID    string `json:"task_id"`
	TaskTime  string `json:"task_time"`
	TaskName  string `json:"task_name"`
	TaskCheck bool   `json:"task_check"`
}

var tasks []Tasks

func allTasks() {
	task := Tasks{
		TaskID:    "1",
		TaskTime:  "21:00",
		TaskName:  "Task1",
		TaskCheck: false,
	}
	tasks = append(tasks, task)
	task2 := Tasks{
		TaskID:    "2",
		TaskTime:  "22:00",
		TaskName:  "Task2",
		TaskCheck: false,
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
		if taskID["id"] == tasks[i].TaskID {
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Tasks
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.TaskID = strconv.Itoa(rand.Intn(1000))
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag := false
	for index, item := range tasks {
		if item.TaskID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			flag = true
			json.NewEncoder(w).Encode(map[string]string{"status": "Success"})
			return
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag := false
	for index, item := range tasks {
		if item.TaskID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Tasks
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.TaskID = params["id"]
			tasks = append(tasks, task)
			flag = true
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	if flag == false {
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
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
    if err != nil {
        fmt.Println("error")
    }

    dbg := os.Getenv("DB")
    connectionPass := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=require",os.Getenv("HOST"),os.Getenv("PORT"),os.Getenv("DBNAME"),os.Getenv("USER"),os.Getenv("PASSWORD"))

    fmt.Printf("DB: %s\n", dbg)
    fmt.Printf("CONNECTIONPASS: %s\n", connectionPass)
	
	db, error := sql.Open(dbg, connectionPass)
	if error != nil {
		log.Fatalln("接続失敗", error)
	}
	defer db.Close()
	if error := db.Ping(); error != nil {
		log.Fatal("PingError: ", err)
	}
	// allTasks()
	// fmt.Println("Hello World")
	// handleRoutes()
}
