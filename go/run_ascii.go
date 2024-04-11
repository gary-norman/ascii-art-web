package ascii_art_web

import (
	"fmt"
	"strings"
)

func RunAscii(input, style, colour, toColour, output, align, reverse string, charMap map[int][]string) string {
	fmt.Println("Entering RunAscii...")
	fmt.Println("------------------------------------------------")
	fmt.Println("input is:", input)
	fmt.Println("style is:", style)
	fmt.Println("colour is:", colour)
	fmt.Println("toColour is:", toColour)
	fmt.Println("output is:", output)
	fmt.Println("align is:", align)

	var word = input
	var words, source []string
	var asciiMap map[int][]string
	//standardMap, shadowMap, thinkertoyMap map[int][]string

	////initialise flags
	//colour := flag.String("color", "default", "colour of the text")
	//output := flag.String("output", "default", "file name which will contain the output of this program")
	//align := flag.String("align", "default", "type of alignment for printed text")
	//reverse := flag.String("reverse", "default", "print text present in the provided ascii art text file")
	//flag.Parse()
	//
	////set all arguments after the flags to otherArgs
	//otherArgs := flag.Args()
	//
	////check flags
	//if !CheckFlagsAndArgs(*colour, *output, *reverse, *align) {
	//	return
	//}
	//
	//prepare correct ascii map(s)
	//if reverse != "default" {

	//	source = ArtFromFile(reverse)
	//} else {
	//prepare the text file for the characters
	source = PrepareBanner(style)

	//if file is non-existent, return
	if source == nil {
		return ""
	}

	//prepare ascii map
	asciiMap = AsciiMap(source)
	fmt.Println(asciiMap)
	//}

	////handle all flags and prepare arguments and variables
	//if colour != "default" {
	//	if toColour == "default" {
	//		toColour = word
	//	}
	//} else
	//if reverse != "default" {
	//	// fmt.Println("source is:", source)

	//}

	//if the word has \n in it, split into separate words
	if word == "" {
		return ""
	} else if strings.Contains(word, "\\n") {
		words = strings.Split(word, "\\n")
	} else {
		words = append(words, word)
	}

	//print words
	fmt.Println("------------------------------------------------")
	fmt.Println("Entering AsciiWords...")
	return AsciiWords(words, asciiMap, charMap, colour, toColour, output, align)

}
