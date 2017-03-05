package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

type pageData struct {
	Date string
	Time string
}

func main() {
	http.HandleFunc("/", renderHomepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderHomepage(wr http.ResponseWriter, r *http.Request) {
	now := time.Now()
	data := pageData{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	renderTemplate(wr, "index.html", data)
}

func renderTemplate(wr io.Writer, templateName string, data interface{}) {
	t, err := template.ParseFiles(templateName)
	if err != nil {
		log.Fatalln("template parsing error: ", err)
	}
	err = t.Execute(wr, data)
	if err != nil {
		log.Fatalln("template executing error: ", err)
	}
}
