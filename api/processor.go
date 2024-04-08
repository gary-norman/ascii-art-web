package api

import (
	ascii_art_web "ascii_art_web/go"
	"html/template"
	"net/http"
)

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
	renderTemplate(w, "index.html", d)
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
