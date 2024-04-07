package main

import (
	"html/template"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		FromArt string
		ToArt   string
	}{
		FromArt: "fromart",
		ToArt:   "toart",
	}
	err = t.Execute(w, data)
	if err != nil {
		return
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}

//func methodHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "GET":
//		fmt.Fprintf(w, "GET request received")
//	case "POST":
//		fmt.Fprintf(w, "POST request received")
//	default:
//		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
//	}
//}

func main() {
	handleRequests()

}
