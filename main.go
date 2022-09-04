package main

import (
	"log"
	"net/http"

	"github.com/nitinjangam/todogo-gin/app"
)

func main() {
	router := app.NewRouter()
	if err := http.ListenAndServe("localhost:9090", router); err != nil {
		log.Fatal(err)
	}
}

func init() {
	app.Tasksdb.Tasks = make(app.Tasks)
}
