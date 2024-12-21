package main

import (
	"net/http"
	"text/template"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r * http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "Hello, world!")
}