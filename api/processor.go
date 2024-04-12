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

	inputText := r.FormValue("generate")
	chosenStyle := r.FormValue("banner")
	chosenColor := r.FormValue("colors")
	chosenAlign := r.FormValue("text-align")
	colorWord := r.FormValue("colour-text")
	var colSlice []rune
	artOutput := ""
	outputResult := "<pre>" + pkg.MakeArt(inputText, pkg.GetChars(pkg.PrepareBan(chosenStyle))) + "</pre>"
	artToText := "Your Art:"
	if ascii_art_web.IsFilePresent(w, r) {
		artOutput = pkg.Reverse("filetoart/" + Reverse(w, r))
		artToText = "Your art says: " + artOutput
		outputResult = "<pre>" + pkg.MakeArt(artOutput, pkg.GetChars(pkg.PrepareBan(chosenStyle))) + "</pre>"
	}
	if chosenColor != "" {
		if colorWord != "" {
			colSlice := []rune(colorWord)
			outputResult = pkg.MakeArtColorized(inputText, pkg.GetChars(pkg.PrepareBan(chosenStyle)), colSlice, chosenColor, false)
		} else {
			outputResult = pkg.MakeArtColorized(inputText, pkg.GetChars(pkg.PrepareBan(chosenStyle)), colSlice, chosenColor, true)
		}
	}
	if chosenAlign == "justify" {
		outputResult = pkg.MakeArtJustified(inputText, pkg.GetChars(pkg.PrepareBan(chosenStyle)))
	}
	d := struct {
		InputText   string
		ChosenStyle string
		ChosenColor string
		ColorWord   string
		ChosenAlign string
		ArtToText   string
		TextToArt   template.HTML
	}{
		InputText:   inputText,
		ChosenStyle: chosenStyle,
		ChosenColor: chosenColor,
		ColorWord:   colorWord,
		ChosenAlign: chosenAlign,
		ArtToText:   artToText,
		TextToArt:   template.HTML(outputResult),
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
