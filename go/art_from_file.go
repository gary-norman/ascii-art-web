package ascii_art_web

import (
	"bufio"
	"log"
	"os"
)

func ArtFromFile(name string) []string {
	var scanned []string
	// Read the ASCII art from a file
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("Error reading ASCII art file:", err)
	}

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		scanned = append(scanned, scan.Text())
	}

	defer file.Close()
	// fmt.Println("scanned:", scanned)
	// Return string
	return scanned
}
