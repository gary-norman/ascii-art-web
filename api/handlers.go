package api

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (se StatusError) Error() string {
	return se.Err.Error()
}

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

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

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	fmt.Println("handling problem :)")
	w.WriteHeader(status)
	//w.Write([]byte(status))
	//http.Error(w, s, status)
	//http.Redirect(w, r, "templates/"+strconv.Itoa(status)+".html", status)
	fmt.Println("redirecting to", strconv.Itoa(status)+".html")
	t, err := template.ParseFiles("templates/" + strconv.Itoa(status) + ".html")
	if err != nil {
		fmt.Println("Error parsing files:", err.Error())
		open500(w)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	return
}
