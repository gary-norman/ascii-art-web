package api

import (
	ascii_art_web "ascii_art_web/go"
	"ascii_art_web/pkg"
	"html/template"
	"net/http"
)

func Processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	chosenInput := r.FormValue("generate")
	chosenStyle := r.FormValue("banner")
	chosenColor := r.FormValue("colors")
	colorInput := r.FormValue("colour-text")
	defaultValue := "default"
	artInput := ascii_art_web.ArtFromFile(w, r)
	chosenAlign := r.FormValue("text-align")
	outputResult := pkg.MakeArt(chosenInput, pkg.GetChars(pkg.PrepareBan(chosenStyle)))
	artOutput := ascii_art_web.CheckReverse(w, r)
	d := struct {
		InputText  string
		Style      string
		Color      string
		ColorWord  string
		FileWant   string
		InputAlign string
		CurrentArt string
		ArtToText  string
		TextToArt  string
	}{
		InputText:  chosenInput,
		Style:      chosenStyle,
		Color:      chosenColor,
		FileWant:   defaultValue,
		ColorWord:  colorInput,
		InputAlign: chosenAlign,
		CurrentArt: artInput,
		ArtToText:  artOutput,
		TextToArt:  outputResult,
	}
	tpl.ExecuteTemplate(w, "index.html", d)
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
