package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", result)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func result(w http.ResponseWriter, r *http.Request) {
	log.Fatal(tpl.ExecuteTemplate(w, "result.html", nil))
}
