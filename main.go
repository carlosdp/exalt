package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getPage)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func getPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index.html").ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
