package main

import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

type Todo struct {
	Title string
	Done bool
}

type ToDoPageData struct {
	PageTitle string
	Todos []Todo
}

func main() {
	r := mux.NewRouter()

	templ := template.Must(template.ParseFiles("template/html/index.html"))
	r.HandleFunc("/todos/{title}", func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]

		data := ToDoPageData{
			PageTitle: title,
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: false},
				{Title: "Task 3", Done: false},
			},
		}
		templ.Execute(w, data)
	})

	http.ListenAndServe(":4040", r)
}