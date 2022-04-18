package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"todo/backend/internal/configs"
	"todo/backend/internal/models"
	"todo/backend/internal/repository"
)

type TaskI interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	GetAllTasks(w http.ResponseWriter, r *http.Request)
	GetTaskById(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	SetTaskStatus(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
}

type Task struct {
	App  *configs.AppConfig
	repo repository.Repository
}

func NewTask(app *configs.AppConfig, repo repository.Repository) *Task {
	return &Task{
		App:  app,
		repo: repo,
	}
}
func (t Task) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.repo.GetAllTasks()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(&tasks)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (t Task) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task *models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(task)
	_, id, err := t.repo.CreateTask(task)
	if err != nil {
		fmt.Println(err)
	}
	message := make(map[string]interface{})
	message["taskId"] = id
	err = json.NewEncoder(w).Encode(&message)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (t Task) GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		return
	}
	driver, err := t.repo.GetTaskById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(&driver)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (t Task) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Println(err)
	}
	updatedTask, err := t.repo.UpdateTask(task)
	if err != nil {
		fmt.Println(err)
	}
	err = json.NewEncoder(w).Encode(&updatedTask)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)

}
func (t Task) SetTaskStatus(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Println(err)
	}
	err = t.repo.SetTaskStatus(task.ID, task.Status)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)

}

func (t Task) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.repo.DeleteTask(id)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)
}
