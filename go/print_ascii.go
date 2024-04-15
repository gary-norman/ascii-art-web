package ascii_art_web

import (
	"fmt"
	"strings"
)

// Prints each ascii character, 1 by 1, line by line and handles colour
func PrintAscii(asciiMap map[int][]string, word string, colour string, toColour string, output string, align string) string {
	var newString, outputString string

	// Iterate over each line of the ASCII art
	for i := 0; i < 8; i++ {
		wordCount := 0
		var sep []string
		var alignedString string
		// Iterate over each character in the line
		for letterNum := 0; letterNum < len(word); letterNum++ {
			// Check if the ASCII value exists in the map and if it matches the current character in the word
			if asciiStr, ok := asciiMap[int(word[letterNum])]; ok {
				//check if there are any characters to colour and render them in the terminal
				if strings.Contains(toColour, string(word[letterNum])) {
					outputString += PrintColorised(colour, asciiStr[i])
					// Print the corresponding line of ASCII art from the map
				} else {

					alignedString += asciiStr[i]
					outputString += asciiStr[i]
					if align != "justify" {
						// fmt.Print("b")
						fmt.Print(asciiStr[i])
					} else {
						if letterNum == 0 {
							sep = append(sep, "<span>"+asciiStr[i])
							fmt.Println("sep is:", sep)
						} else {
							if word[letterNum] == ' ' {
								sep[wordCount] += "</span>"
								wordCount++
								sep = append(sep, "<span>"+asciiStr[i])
								fmt.Println("*sep is:", sep)
							} else {
								sep[wordCount] += (asciiStr[i] + "</span>")
								fmt.Println("**sep is:", sep)
							}
						}
					}
				}
			}
		}
		outputString += "\n"
		if align != "justify" {
			newString = outputString
		} else {
			fmt.Println("------------------------------------------------")
			fmt.Println("Entering AlignJustified...")
			newString = AlignJustified(align, sep, alignedString)
			fmt.Println("Aligned newString is:", newString)
		}

	}
	return newString
	////handle outputfile if present
	//if output != "default" {
	//	MakeOutput(output, outputString)
	//}
}
