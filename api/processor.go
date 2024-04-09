package api

import (
	"ascii_art_web/pkg"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func reverse(r *http.Request) string {
	// Limit the size of the incoming file to prevent memory issues
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return ""
	} // Limit upload size to 10MB

	// Retrieve the file from form data
	file, handler, err := r.FormFile("file-drop")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return ""
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Create a file in the server's local storage
	dst, err := os.Create("filetoart/" + handler.Filename)
	if err != nil {
		fmt.Println("Error Creating the File")
		fmt.Println(err)
		return ""
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {

		}
	}(dst)

	// Copy the uploaded file data to the server's file
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println("Error Saving the File")
		fmt.Println(err)
		return ""
	}
	return handler.Filename
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	chosenInput := r.FormValue("generate")
	chosenStyle := r.FormValue("banner")
	chosenColor := r.FormValue("colors")
	colorInput := r.FormValue("colour-text")
	defaultValue := "default"
	//artInput := r.FormValue("file-drop")
	artInOut := pkg.Reverse("filetoart/" + reverse(r))
	chosenAlign := r.FormValue("text-align")
	//outputResult := "poo"
	outputResult := pkg.MakeArt(chosenInput, pkg.GetChars(pkg.PrepareBan(chosenStyle)))
	//fmt.Print(chosenInput)
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
		ArtToText:  artInOut,
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
