package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

func main() {
	todos := map[string][]Todo{
		"Todos": {
			{Title: "Todo 1", Done: false},
			{Title: "Todo 2", Done: true},
			{Title: "Todo 3", Done: false},
		},
	}

	// Serve index.html
	h1 := func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, todos)

	}

	// Handle Post function
	h2 := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		title := r.PostFormValue("title")
		todos["Todos"] = append(todos["Todos"], Todo{Title: title, Done: false})

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "todo-items", todos)

	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-todo/", h2)

	//Serve
	log.Fatal(http.ListenAndServe(":8080", nil))
}
