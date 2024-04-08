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
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("icons"))))
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
	input := r.FormValue("input")
	style := r.FormValue("style")
	color := r.FormValue("colors")
	colorInput := r.FormValue("colour-text")
	defaultValue := "default"
	artInput := r.FormValue("art-input")
	align := r.FormValue("text-align")

	d := struct {
		InputText  string
		Style      string
		Color      string
		ColorWord  string
		FileWant   string
		InputAlign string
		InputArt   string
		Output     string
	}{
		InputText:  input,
		Style:      style,
		Color:      color,
		FileWant:   defaultValue,
		ColorWord:  colorInput,
		InputAlign: align,
		InputArt:   artInput,
		Output:     ascii_art_web.RunAscii(input, color, colorInput, defaultValue, align, artInput),
	}
	renderTemplate(w, "html.html", d)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Parse the template file
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Error parsing template file", http.StatusInternalServerError)
		return
	}

	// Execute the template with the provided data and write the result to the response writer
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
