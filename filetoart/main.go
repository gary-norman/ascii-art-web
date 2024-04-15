package main

import (
	"flag"
	"piscine"
	"strings"
)

func main() {
	var word, toColour string
	var words, source []string
	var asciiMap, standardMap, shadowMap, thinkertoyMap map[int][]string

	//initialise flags
	colour := flag.String("color", "default", "colour of the text")
	output := flag.String("output", "default", "file name which will contain the output of this program")
	align := flag.String("align", "default", "type of alignment for printed text")
	reverse := flag.String("reverse", "default", "print text present in the provided ascii art text file")
	flag.Parse()

	//set all arguments after the flags to otherArgs
	otherArgs := flag.Args()

	//check flags
	if !piscine.CheckFlagsAndArgs(*colour, *output, *reverse, *align) {
		return
	}

	//prepare correct ascii map(s)
	if *reverse != "default" {
		standardMap = piscine.AsciiMap(piscine.PrepareBanner("standard"))
		shadowMap = piscine.AsciiMap(piscine.PrepareBanner("shadow"))
		thinkertoyMap = piscine.AsciiMap(piscine.PrepareBanner("thinkertoy"))
		source = piscine.ArtFromFile(*reverse)
	} else {
		//prepare the text file for the characters
		source = piscine.PrepareBanner("")
		//if file is non-existent, return
		if source == nil {
			return
		}

		//prepare ascii map
		asciiMap = piscine.AsciiMap(source)
	}

	//handle all flags and prepare arguments and variables
	if *colour != "default" {
		if len(otherArgs) == 1 {
			word = otherArgs[0]
			toColour = word
		} else {
			toColour = otherArgs[0]
			word = otherArgs[1]
		}
	} else if *reverse != "default" {
		// fmt.Println("source is:", source)
		emptyCols := piscine.RemoveValidSpaceIndex(piscine.GetEmptyCols(source))
		charMap := piscine.CharMap(piscine.ArtToSingleLine(source), emptyCols)
		piscine.AsciiToChars(charMap, standardMap, shadowMap, thinkertoyMap)
	} else {
		word = otherArgs[0]
	}

	//if the word has \n in it, split into separate words
	if word == "" {
		return
	} else if strings.Contains(word, "\\n") {
		words = strings.Split(word, "\\n")
	} else {
		words = append(words, word)
	}

	//print words
	piscine.PrintAsciiWords(words, asciiMap, *colour, toColour, *output, *align)

}
