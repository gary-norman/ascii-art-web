package ascii_art_web

// PrintAsciiWords Prints all words inside a string, that were and weren't separated by \n
func PrintAsciiWords(words []string, asciiMap map[int][]string, colour string, toColour string, outputFile string, align string) string {
	var newString string
	//loop over words array to print each word
	for i := 0; i < len(words); i++ {
		for words[i] == "" {
			words = append(words[:i], words[i+1:]...)
		}

		newString = PrintAscii(asciiMap, words[i], colour, toColour, outputFile, align)

	}
	return newString
}
