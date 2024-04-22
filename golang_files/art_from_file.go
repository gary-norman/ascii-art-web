package asciiartweb

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ArtFromFile(w http.ResponseWriter, r *http.Request) string {
	fmt.Println("entering ArtFromFile")

	var value string

	// Limit the size of the incoming file to prevent memory issues
	err := r.ParseMultipartForm(10 << 20) // Limit upload size to 10MB
	if err != nil {
		fmt.Println("Error parsing form data:", err)
		return ""
	}

	// Retrieve the file from form data
	file, handler, err := r.FormFile("file-drop")
	if err != nil {
		fmt.Println("Error retrieving the file:", err)
		return ""
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing the file:", err)
		}
	}(file)

	// Create a file in the server's local storage
	dst, err := os.Create(filepath.Join("filetoart", handler.Filename))
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return ""
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			fmt.Println("Error closing the file:", err)
		}
	}(dst)

	// Copy the uploaded file data to the server's file
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println("Error saving the file:", err)
		return ""
	}

	// Read the content of the text file
	content, err := os.ReadFile(filepath.Join("filetoart", handler.Filename))
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return ""
	}

	// Convert byte slice to string
	value = string(content)

	fmt.Println("Value extracted from file:", value)
	return value
}

func ArtFromFileLines(w http.ResponseWriter, r *http.Request) []string {
	fmt.Println("entering ArtFromFileLines")
	fmt.Println("------------------------------------------------")
	// Limit the size of the incoming file to prevent memory issues
	err := r.ParseForm() // Limit upload size to 10MB
	if err != nil {
		fmt.Println("Error parsing form data:", err)
	}

	fmt.Println("no err")
	// Retrieve the file from form data
	file, handler, err := r.FormFile("file-drop")
	if err != nil {
		fmt.Println("Error retrieving the file:", err)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing the file:", err)
		}
	}(file)

	// Create a file in the server's local storage
	dst, err := os.Create(filepath.Join("filetoart", handler.Filename))
	if err != nil {
		fmt.Println("Error creating the file:", err)
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			fmt.Println("Error closing the file:", err)
		}
	}(dst)

	// Copy the uploaded file data to the server's file
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println("Error saving the file:", err)
	}
	// Read the content of the text file
	content, err := os.ReadFile(filepath.Join("filetoart", handler.Filename))
	if err != nil {
		fmt.Println("Error reading the file:", err)
	}
	newFile := strings.Split(string(content), "\n")

	defer file.Close()
	// Return string
	return newFile
}
