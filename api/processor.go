package api

import (
	asciiartweb "asciiartweb/golang_files"
	"asciiartweb/pkg"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Processor(w http.ResponseWriter, r *http.Request) {
	d := struct {
		InputText   string
		ChosenStyle string
		ChosenColor string
		ColorWord   string
		ChosenAlign string
		ArtToText   string
		TextToArt   template.HTML
	}{}

	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	inputText := r.FormValue("generate")
	chosenStyle := r.FormValue("banner")
	chosenColor := r.FormValue("colors")
	chosenAlign := r.FormValue("text-align")
	colorWord := r.FormValue("colour-text")
	var colSlice []rune
	artOutput := ""
	outputResult := ""
	//outputResult := "<pre>" + pkg.MakeArt(inputText, pkg.GetChars(pkg.PrepareBan(chosenStyle))) + "</pre>"
	outputResult, chosenAlign = pkg.MakeArtJustified(inputText, pkg.GetChars(pkg.PrepareBan(chosenStyle)), chosenAlign)
	artToText := "Your Art:"
	// justify alignment
	//if chosenAlign == "justify" {
	//	outputResult, chosenAlign = pkg.MakeArtJustified(inputText, pkg.GetChars(pkg.PrepareBan(chosenStyle)), chosenAlign)
	//
	//}
	//colourise art
	if chosenColor != "" {
		if colorWord != "" {
			colSlice := []rune(colorWord)
			outputResult = pkg.MakeArtColorized(inputText, pkg.GetChars(pkg.PrepareBan(chosenStyle)), colSlice, chosenColor, false)
		} else {
			outputResult = pkg.MakeArtColorized(inputText, pkg.GetChars(pkg.PrepareBan(chosenStyle)), colSlice, chosenColor, true)
		}
	}
	// reverse lookup
	if asciiartweb.IsFilePresent(w, r) {
		fmt.Println("testing - file present: ", GetFileName(w, r))
		artOutput = pkg.Reverse("filetoart/" + GetFileName(w, r))
		artToText = "Your art says: " + artOutput
		outputResult = "<pre>" + asciiartweb.ArtFromFile(w, r) + "</pre>"
	}

	// write art to file for download
	err := os.WriteFile("arttofile/yourart.txt", []byte(pkg.MakeArt(inputText, pkg.GetChars(pkg.PrepareBan(chosenStyle)))+"\n"), 0644)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return // Exit the program on error
	}
	// force 400 error
	if inputText == "" && !asciiartweb.IsFilePresent(w, r) {
		fmt.Println("Error1 in Processor")
		ErrorHandler(w, r, http.StatusBadRequest)
	} else if inputText == "" && asciiartweb.IsFilePresent(w, r) {

	} else if inputText != "" && colorWord != "" && chosenColor == "" {
		fmt.Println("Error2 in Processor")
		ErrorHandler(w, r, http.StatusBadRequest)
	}

	d.InputText = inputText
	d.ChosenStyle = chosenStyle
	d.ChosenColor = chosenColor
	d.ColorWord = colorWord
	d.ChosenAlign = chosenAlign
	d.ArtToText = artToText
	d.TextToArt = template.HTML(outputResult)

	//if all good, status 200, writing it to head would make it redundant, as per  " http: superfluous response.WriteHeader call from ascii_art_web/api.Processor (processor.golang_files:87)"
	err = tpl.ExecuteTemplate(w, "result.html", d)
	if err != nil {
		fmt.Println("Error is:", err)
		var e Error
		switch {
		case errors.As(err, &e):
			fmt.Println("Error3 in Processor")
			ErrorHandler(w, r, e.Status())
		default:
			fmt.Println("Error4 in Processor")
			ErrorHandler(w, r, http.StatusInternalServerError)
		}
		return
	}
}

// force 500 error
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
