package api

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func HandleRequestsGary() {
	http.HandleFunc("/", HomePageGary)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("icons"))))
	http.Handle("/ascii_styles/", http.StripPrefix("/ascii_styles/", http.FileServer(http.Dir("ascii_styles"))))
	http.HandleFunc("/process", Processor)
	//http.HandleFunc("/upload" Reverse)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
