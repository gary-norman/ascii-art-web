package ascii_art_web

import "fmt"

// AsciiWords Prints all words inside a string, that were and weren't separated by \n
func AsciiWords(words []string, asciiMap, charMap map[int][]string, colour string, toColour string, outputFile string, align string) string {
	fmt.Println("------------------------------------------------")

	var newString string
	//loop over words array to print each word
	for i := 0; i < len(words); i++ {
		for words[i] == "" {
			words = append(words[:i], words[i+1:]...)
		}
		if align == "left" {
			fmt.Println("Entering PrintAscii...")
			newString += PrintAscii(asciiMap, words[i], colour, toColour, outputFile, align)
		} else {
			AlignWords(charMap, words, align)
		}

	}
	return newString
}
