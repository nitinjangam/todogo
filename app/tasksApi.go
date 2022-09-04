package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Task struct {
	Description string `json:"description"`
	TaskID      string `json:"task-id"`
	Completed   bool   `json:"completed"`
}

type Tasks map[string]Task

type TasksDB struct {
	Tasks Tasks `json:"tasks"`
}

var Tasksdb TasksDB

type Response struct {
	Tasks Tasks  `json:"tasks"`
	Error string `json:"error"`
}

func getAllTasks(ctx *gin.Context) {
	response := Response{}
	response.Tasks = make(Tasks)
	if len(Tasksdb.Tasks) == 0 {
		response.Error = "No tasks in db"
		ctx.IndentedJSON(http.StatusNotFound, &response)
		return
	}
	response.Tasks = Tasksdb.Tasks
	ctx.IndentedJSON(http.StatusOK, &response)
}

func getTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")
	response := Response{}
	if _, ok := Tasksdb.Tasks[id]; !ok {
		response.Error = "Task not found"
		ctx.AbortWithStatusJSON(http.StatusNotFound, &response)
		return
	}
	response.Tasks = make(Tasks)
	response.Tasks[id] = Tasksdb.Tasks[id]
	ctx.IndentedJSON(http.StatusOK, &response)
}

func addTask(ctx *gin.Context) {
	response := Response{}
	response.Tasks = make(Tasks)
	newTask := Task{}
	if err := ctx.BindJSON(&newTask); err != nil {
		response.Error = "one or more fields missing"
		ctx.IndentedJSON(http.StatusBadRequest, &response)
	}
	id := uuid.New()
	newTask.TaskID = id.String()
	Tasksdb.Tasks[id.String()] = newTask
	response.Tasks[id.String()] = newTask
	ctx.IndentedJSON(http.StatusCreated, &response)
}

func modifyTask(ctx *gin.Context) {
	response := Response{}
	response.Tasks = make(Tasks)
	id := ctx.Param("id")
	if _, ok := Tasksdb.Tasks[id]; !ok {
		response.Error = "Task not found"
		ctx.AbortWithStatusJSON(http.StatusNotFound, &response)
		return
	}
	updatedTask := Task{}
	if err := ctx.BindJSON(&updatedTask); err != nil {
		response.Error = "one or more fields missing"
		ctx.IndentedJSON(http.StatusBadRequest, &response)
	}
	updatedTask.TaskID = Tasksdb.Tasks[id].TaskID
	updatedTask.Description = Tasksdb.Tasks[id].Description
	Tasksdb.Tasks[id] = updatedTask
	response.Tasks[id] = updatedTask
	ctx.IndentedJSON(http.StatusCreated, &response)
}

func deleteTask(ctx *gin.Context) {
	response := Response{}
	id := ctx.Param("id")
	if _, ok := Tasksdb.Tasks[id]; !ok {
		response.Error = "Task not found"
		ctx.AbortWithStatusJSON(http.StatusNotFound, &response)
		return
	}
	delete(Tasksdb.Tasks, id)
	ctx.IndentedJSON(http.StatusAccepted, "")
}
