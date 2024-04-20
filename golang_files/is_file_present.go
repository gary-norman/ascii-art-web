package asciiartweb

import (
	"fmt"
	"mime/multipart"
	"net/http"
)

func IsFilePresent(w http.ResponseWriter, r *http.Request) bool {

	// Retrieve the file from form data
	file, handler, err := r.FormFile("file-drop")
	if err != nil {
		fmt.Println("No file in Reverse_kamil")
		//fmt.Println(err)
		return false
	}
	fmt.Println("file is:", file)
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	if handler.Filename != "" {
		return true
	}

	return false
}
