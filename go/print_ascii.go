package ascii_art_web

import (
	"fmt"
	"strings"
)

var outputString string

// Prints each ascii character, 1 by 1, line by line and handles colour
func PrintAscii(asciiMap map[int][]string, word string, colour string, toColour string, output string, align string) string {
	var newString string
	spaces := 0
	// Iterate over each line of the ASCII art
	for i := 0; i < 8; i++ {
		wordCount := 0
		var sep []string
		var alignedString string

		// Iterate over each character in the line
		for letterNum := 0; letterNum < len(word); letterNum++ {

			// Check if the ASCII value exists in the map and if it matches the current character in the word
			if asciiStr, ok := asciiMap[int(word[letterNum])]; ok {

				//check if there are any characters to colour
				if strings.Contains(toColour, string(word[letterNum])) {
					outputString += asciiStr[i]
					PrintColorised(colour, asciiStr[i])
					// Print the corresponding line of ASCII art from the map
				} else {

					alignedString += asciiStr[i]

					outputString += asciiStr[i]

					if align == "default" {
						// fmt.Print("b")
						fmt.Print(asciiStr[i])
					} else {
						spaces = len(word) - 1

						if letterNum == 0 {
							sep = append(sep, asciiStr[i])
						} else {
							if word[letterNum] == ' ' {
								wordCount++
								sep = append(sep, asciiStr[i])
							} else {
								sep[wordCount] += asciiStr[i]
							}
						}

					}
				}
			}
		}
		outputString += "\n"

		if align == "default" {
			fmt.Println()
		} else {
			PrintAligned(align, sep, spaces, alignedString)
		}
		newString = outputString
	}
	fmt.Println("testing 2")
	return newString
	////handle outputfile if present
	//if output != "default" {
	//	MakeOutput(output, outputString)
	//}
}