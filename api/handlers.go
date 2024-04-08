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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Limit the size of the incoming file to prevent memory issues
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return
	} // Limit upload size to 10MB

	// Retrieve the file from form data
	file, handler, err := r.FormFile("file-drop")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
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
		return
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
		return
	}

	artInput := r.FormValue("file-drop")
	artInOut := pkg.Reverse("/filetoart/" + artInput)
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
		//InputText:  chosenInput,
		//Style:      chosenStyle,
		//Color:      chosenColor,
		//FileWant:   defaultValue,
		//ColorWord:  colorInput,
		//InputAlign: chosenAlign,
		ArtToText: artInOut,
		//TextToArt:  outputResult,
	}
	tpl.ExecuteTemplate(w, "index.html", d)
}

func HandleRequestsGary() {
	http.HandleFunc("/", HomePageGary)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("icons"))))
	//http.Handle("/pkg/", http.StripPrefix("/pkg/", http.FileServer(http.Dir("pkg"))))
	http.Handle("/ascii_styles/", http.StripPrefix("/ascii_styles/", http.FileServer(http.Dir("ascii_styles"))))
	http.HandleFunc("/process", processor)
	http.HandleFunc("/upload", uploadFileHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
