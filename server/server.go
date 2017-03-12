package main

import (
	"html/template"
	"net/http"
	"time"
)

type page struct {
	Title string
	Date  string
}

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", renderHomepage)
	http.ListenAndServe(":8080", nil)
}

func renderHomepage(w http.ResponseWriter, r *http.Request) {
	page := page{"Hello Go", time.Now().Format(time.RFC1123)}
	renderTemplate(w, "index.html", page)
}

func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	err := templates.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
