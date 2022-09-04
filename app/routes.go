package app

import "github.com/gin-gonic/gin"

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func NewRouter() *gin.Engine {
	router := gin.Default()

	for _, v := range routes {
		router.Handle(v.Method, v.Pattern, v.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"GET",
		"/task",
		getAllTasks,
	},
	Route{
		"GET",
		"/task/:id",
		getTaskByID,
	},
	Route{
		"POST",
		"/task",
		addTask,
	},
	Route{
		"PUT",
		"/task/:id",
		modifyTask,
	},
	Route{
		"DELETE",
		"/task/:id",
		deleteTask,
	},
}
