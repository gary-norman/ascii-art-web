package ascii_art_web

import (
	"fmt"
	"os"
)

func MakeOutput(outputFile string, result string) {
	// save content to new file
	errout := os.WriteFile(outputFile, []byte(result), 0666)
	if errout != nil {
		fmt.Println(errout)
		return
	}
}
