package api

import (
	ascii_art_web "ascii_art_web/go"
	"ascii_art_web/pkg"
	"errors"
	"fmt"
	"html/template"
	"net/http"
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
		ErrorHandler(w, r, http.StatusBadGateway)
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
	fmt.Println("r.URL.Path", r.URL.Path)
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

	if inputText == "" {
		//http.Error(w, "400 - Bad Request: Missing form fields", http.StatusBadRequest)
		fmt.Println("Error1 in Processor")
		ErrorHandler(w, r, http.StatusBadRequest)
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

	//if all good, status 200, writing it to head would make it redundant, as per  " http: superfluous response.WriteHeader call from ascii_art_web/api.Processor (processor.go:87)"
	err := tpl.ExecuteTemplate(w, "result.html", d)
	if err != nil {
		fmt.Println("err is:", err)
		var e Error
		switch {
		case errors.As(err, &e):
			//http.Error(w, e.Error(), e.Status())
			fmt.Println("Error3 in Processor")
			ErrorHandler(w, r, e.Status())
		default:
			fmt.Println("Error4 in Processor")
			ErrorHandler(w, r, http.StatusInternalServerError)
		}
		return
	}
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

func open500(w http.ResponseWriter) {
	w.WriteHeader(500)
	t, err := template.ParseFiles("templates/500.html")
	if err != nil {
		fmt.Println("Error parsing files:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
}
