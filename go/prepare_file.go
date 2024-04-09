package ascii_art_web

import (
	"bufio"
	"fmt"
	"os"
)

func PrepareFile(style string) []string {
	var scanned []string
	file, err := os.Open("ascii_styles/" + style + ".txt")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		scanned = append(scanned, scan.Text())
	}

	defer file.Close()
	return scanned
}
