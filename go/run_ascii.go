package ascii_art_web

import (
	"fmt"
	"strings"
)

func RunAscii(input, colour, toColour, output, align, reverse string) string {
	var word = input
	var words, source []string
	var asciiMap, standardMap, shadowMap, thinkertoyMap map[int][]string
	var placeholderVar string

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
	//check flags
	if !CheckFlagsAndArgs(colour, output, reverse, align) {
		return ""
	}
	//
	//prepare correct ascii map(s)
	if reverse != "default" {
		//standardMap = AsciiMap(PrepareBanner("standard"))
		//shadowMap = AsciiMap(PrepareBanner("shadow"))
		//thinkertoyMap = AsciiMap(PrepareBanner("thinkertoy"))
		source = ArtFromFile(reverse)
	} else {
		//prepare the text file for the characters
		source = PrepareBanner("")
		//if file is non-existent, return
		if source == nil {
			return ""
		}

		//prepare ascii map
		asciiMap = AsciiMap(source)
	}

	//handle all flags and prepare arguments and variables
	if colour != "default" {
		if len(placeholderVar) == 1 {
			word = placeholderVar
			toColour = word
		} else {
			toColour = placeholderVar
			word = placeholderVar
		}
	} else if reverse != "default" {
		// fmt.Println("source is:", source)
		emptyCols := RemoveValidSpaceIndex(GetEmptyCols(source))
		charMap := CharMap(ArtToSingleLine(source), emptyCols)
		AsciiToChars(charMap, standardMap, shadowMap, thinkertoyMap)
	} else {
		word = placeholderVar
	}

	//if the word has \n in it, split into separate words
	if word == "" {
		return ""
	} else if strings.Contains(word, "\\n") {
		words = strings.Split(word, "\\n")
	} else {
		words = append(words, word)
	}

	//print words
	fmt.Println("testing 1")
	return PrintAsciiWords(words, asciiMap, colour, toColour, output, align)

}
