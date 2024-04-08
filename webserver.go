package main

import (
	"html/template"
	"net/http"
	"piscine/pkg"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	type PageData struct {
		FromArt string
		ToArt   string
	}
	data := PageData{
		FromArt: "fromart",
		ToArt:   pkg.MakeArt("test", pkg.GetChars(pkg.PrepareBan("standard"))),
	}
	err = t.Execute(w, data)
	if err != nil {
		return
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("icons"))))
	//http.Handle("/pkg/", http.StripPrefix("/pkg/", http.FileServer(http.Dir("pkg"))))
	http.Handle("/ascii_styles/", http.StripPrefix("/ascii_styles/", http.FileServer(http.Dir("ascii_styles"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

//pkg methodHandler(w http.ResponseWriter, r *http.Request) {
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
