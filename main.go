package main

import (
	ascii_art_web "ascii_art_web/go"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func homePage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		return
	}
}

func handleRequests() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("icons"))))
	http.Handle("/ascii_styles/", http.StripPrefix("/ascii_styles/", http.FileServer(http.Dir("ascii_styles"))))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/process", processor)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	chosenInput := r.FormValue("generate")
	chosenStyle := r.FormValue("style")
	chosenColor := r.FormValue("colors")
	colorInput := r.FormValue("colour-text")
	defaultValue := "default"
	artInput := r.FormValue("file-drop")
	chosenAlign := r.FormValue("text-align")
	outputResult := ascii_art_web.RunAscii(chosenInput, chosenColor, colorInput, defaultValue, chosenAlign, artInput)

	d := struct {
		InputText  string
		Style      string
		Color      string
		ColorWord  string
		FileWant   string
		InputAlign string
		ArtToText  string
		TextToArt  string
	}{
		InputText:  chosenInput,
		Style:      chosenStyle,
		Color:      chosenColor,
		FileWant:   defaultValue,
		ColorWord:  colorInput,
		InputAlign: chosenAlign,
		ArtToText:  artInput,
		TextToArt:  outputResult,
	}
	tpl.ExecuteTemplate(w, "index.html", d)
}
