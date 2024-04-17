package api

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func GetFileName(w http.ResponseWriter, r *http.Request) string {

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
