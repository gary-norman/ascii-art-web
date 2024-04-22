package api

import (
	asciiartweb "asciiartweb/golang_files"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Downloader(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entering Downloader func")
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	fmt.Println("Getting artwork")
	artworkLines := asciiartweb.ArtFromFileLines(w, r)
	//process artwork text lines, add \n at the end of each and join
	fmt.Println("Adding \\n's to artwork")
	artworkText := strings.Join(artworkLines, "\\n")

	fmt.Println("Saving artwork to file")
	// Call the function to save the artwork to a file
	SaveToFile(artworkText)

	fmt.Println("Opening the file")
	// Open the file for reading
	file, err := os.Open("your_artwork.txt")
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Set the appropriate headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=your_artwork.txt")
	w.Header().Set("Content-Type", "text/plain")

	fmt.Println("Copying the file")
	// Copy the file contents to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error copying file", http.StatusInternalServerError)
		return
	}
}

func SaveToFile(artworkText string) {
	// Specify the file name you want to save the artwork to
	fileName := "your_artwork.txt"

	// Write the artwork to the file
	err := os.WriteFile(fileName, []byte(artworkText), 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	fmt.Printf("Artwork saved to %s\n", fileName)
}
