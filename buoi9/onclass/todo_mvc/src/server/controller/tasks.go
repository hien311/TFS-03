package controller

import (
	"encoding/json"
	"net/http"
	"todo_mvc/db"
	"todo_mvc/model"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []model.Task
	db.Connect().Find(&tasks)

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		logrus.Error("Parse data fail ", err)
	}
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		logrus.Error("parse request body: fail ", err)
		json.NewEncoder(w).Encode("fail to add task")
		return
	}

	logrus.Info("Add Task to db")
	db.Connect().Create(&task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	id := mux.Vars(r)["id"]
	logrus.Info("Delete Task id =", id)
	db.Connect().Delete(task, id)
}

func EditTask(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Edit Task")
	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		logrus.Error("parse request body: fail ", err)
		json.NewEncoder(w).Encode("fail to add task")
		return
	}

	db.Connect().Save(task)
}

func CompleteAllTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	logrus.Info("Complete All Task")
	db.Connect().Model(&task).Where("status = ?", false).Update("status", true)
}

func DeleteCompletedTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	logrus.Info("Delete Completed Tasks")
	db.Connect().Delete(task, "status = ?", true)
}
