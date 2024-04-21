package pkg

import (
	"fmt"
	"os"
	"strings"
)

// MakeArt Transform the input text origString to the output art, line by line
func MakeArt(origString string, y map[int][]string) string {
	var art string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")               // split the input into slices
	for _, word := range wordSlice {                                // loop over the word to get the characters
		for j := 0; j < len(y[32]); j++ { // loop over each vertical line of the word
			var line string
			for _, letter := range word { // loop over each character
				line = line + y[int(letter)][j] // add each line of the character to the line string
			}
			art += line + "\n" // add each line string (followed by a line break) to the final output
			line = ""
		}
	}
	art = strings.TrimRight(art, "\n") // remove the final line break
	return art
}

// MakeArtAligned Transform the input text origString to the output art, line by line, with left, right, or center aligned content
func MakeArtAligned(origString string, y map[int][]string, ds []int, ws Winsize, divider int) string {
	var art string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	for i := 0; i < len(wordSlice); i++ {
		for j := 0; j < len(y[32]); j++ {
			var line string
			art += strings.Repeat(" ", (int(ws.Col)-ds[i])/divider)
			for _, letter := range wordSlice[i] {
				line = line + y[int(letter)][j]
			}
			art += line + "\n"
			line = ""
		}
	}
	art = strings.TrimRight(art, "\n")
	return art
}

// prepareInputText prepares the input text for processing (splits into [lines][words]
func prepareInputText(origString string) ([]string, [][]string) {
	// prepare input for processing
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	//split input into lines
	var lines [][]string
	for _, str := range wordSlice {
		splitStr := strings.Split(str, "\n")
		lines = append(lines, splitStr)
	}
	return wordSlice, lines
}

// alignWordsToPre processes the words on (each) line into separate <pre>s
func alignWordsToPre(words []string, y map[int][]string, wordSlice []string, align string) string {
	var art string
	if len(words) > 1 {
		var line string
		for _, word := range words {
			switch align {
			case "left":
				line = "<pre style=\"margin: 0 3.2rem 0 0\">"
			case "right":
				line = "<pre style=\"margin: 0 0 0 3.2rem\">"
			case "center":
				line = "<pre style=\"margin: 0 1.6rem 0 .8rem\">"
			default:
				line = "<pre>"
			}
			//fmt.Printf("Added <pre>\n")
			for j := 0; j < len(y[32]); j++ {
				for _, letter := range word {
					line = line + y[int(letter)][j]
					//fmt.Printf("Added letter #%v\n", j)
				}
				line += "\n"
				//fmt.Printf("Added new line\n")
			}
			line += "</pre>"
			art += line
		}

	} else {
		line := "<pre>"
		for _, word := range wordSlice {
			for j := 0; j < len(y[32]); j++ {
				for _, letter := range word {
					line = line + y[int(letter)][j]
				}
				line += "\n"
			}
		}
		art += line + "</pre>"
		line = "<pre>"
	}
	return art
}

func colorWordsToPre(words []string, y map[int][]string, letters []rune, color string, colorAll bool) string {
	var line string
	var art string
	for _, word := range words {
		if colorAll {
			fmt.Printf("Printing colorAll: %v\n", colorAll)
			for _, letter := range word {
				line = line + "<pre class=" + color + ">"
				for j := 0; j < len(y[32]); j++ {
					line = line + y[int(letter)][j] + "\n"
				}
				line = line + "</pre>"
			}
		} else {
			for _, letter := range word {
				if Contains(letters, letter) {
					fmt.Printf("Printing coloured: %v\n", letter)
					line = line + "<pre class=" + color + ">"
					for j := 0; j < len(y[32]); j++ {
						line = line + y[int(letter)][j] + "\n"
					}
					line = line + "</pre>"
				} else {
					fmt.Printf("Printing non-coloured\n")
					line = line + "<pre>"
					for j := 0; j < len(y[32]); j++ {
						line = line + y[int(letter)][j] + "\n"
					}
					line = line + "</pre>"
				}
			}
		}
		art += line + "\n"
		line = ""
	}
	//art = strings.TrimRight(art, "\n")
	return art
}

// MakeArtJustified Transform the input text origString to the output art, line by line, with justified content
func MakeArtJustified(origString string, y map[int][]string, align string, letters []rune, color string, colorAll bool) (string, string) {
	wordSlice, lines := prepareInputText(origString)
	fmt.Println(wordSlice)
	var art string
	var justification string
	//processor for multiple line
	if len(lines) > 1 {
		for _, newLine := range lines {
			var newWords []string
			for _, word := range newLine {
				newWords = append(newWords, strings.Split(word, " ")...)
			}
			art += "<div class=\"justifiedOutput" + align + "\">"
			if color != "" {
				art += colorWordsToPre(newWords, y, letters, color, colorAll)
			} else {
				art += alignWordsToPre(newWords, y, newLine, align)
			}
			art += "</div>"
			justification = "Multiline" //sets the class of the <pre>
		}
		// processor for single lines
	} else {
		for _, newLine := range lines {
			var newWords []string
			for _, word := range newLine {
				newWords = append(newWords, strings.Split(word, " ")...)
			}
			if color != "" {
				art += colorWordsToPre(newWords, y, letters, color, colorAll)
			} else {
				art += alignWordsToPre(newWords, y, newLine, align)
			}
			justification = align
		}
	}
	return art, justification
}

// MakeArtColorized Transform the input text origString to the output art, line by line, colorizing specified text
func MakeArtColorized(origString string, y map[int][]string, letters []rune, color string, colorAll bool) string {
	wordSlice, lines := prepareInputText(origString)
	fmt.Println(lines)
	var line string
	var art string
	for _, word := range wordSlice {
		if colorAll {
			for _, letter := range word {
				line = line + "<pre class=" + color + ">"
				for j := 0; j < len(y[32]); j++ {
					line = line + y[int(letter)][j] + "\n"
				}
				line = line + "</pre>"
			}
		} else {
			for _, letter := range word {
				if Contains(letters, letter) {
					line = line + "<pre class=" + color + ">"
					for j := 0; j < len(y[32]); j++ {
						line = line + y[int(letter)][j] + "\n"
					}
					line = line + "</pre>"
				} else {
					line = line + "<pre>"
					for j := 0; j < len(y[32]); j++ {
						line = line + y[int(letter)][j] + "\n"
					}
					line = line + "</pre>"
				}
			}
		}
		art += line + "\n"
		line = ""
	}
	//art = strings.TrimRight(art, "\n")
	return art
}

func Reverse(fileName string) string {
	fmt.Print("reverse")
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return "error" // Exit the program on error
	}
	//defer pkg(file os.File) {
	//	err := file.Close()
	//	if err != nil {
	//		fmt.Println("Error closing the file:", err)
	//	}
	//}(file)
	source := FileToVariable(file)
	emptyCols := RemoveValidSPaceIndex(GetEmptyCols(source))
	charMap := GetInputChars(ArtToSingleLine(source), emptyCols)
	mapStandard := GetChars(PrepareBan("standard"))
	mapShadow := GetChars(PrepareBan("shadow"))
	mapThinkertoy := GetChars(PrepareBan("thinkertoy"))
	return AsciiToChars(charMap, mapStandard, mapShadow, mapThinkertoy)

}

//	if *color != "default" {
//		var colored string
//		var colorAll bool
//		var colSLice []rune
//		colorAll = true
//		if len(additionalArgs) == 2 {
//			colorAll = false
//			colored = additionalArgs[0]
//			colSLice = []rune(colored)
//			input = additionalArgs[1]
//		}
//		fmt.Println(MakeArtColorized(input, GetChars(PrepareBan("")), colSLice, *color, colorAll))
//		return
//	}
//	if *output != "default" {
//
//		err := os.WriteFile(*output, []byte(MakeArt(input, GetChars(PrepareBan(bannerStyle)))+"\n"), 0644)
//		if err != nil {
//			fmt.Println("Error writing to the file:", err)
//			return // Exit the program on error
//		}
//

//	if *align == "right" {
//
//		ws := GetWinSize()
//		ds := GetArtWidth(input, GetCharsWidth(PrepareBan(bannerStyle)))
//		fmt.Println(MakeArtAligned(input, GetChars(PrepareBan(bannerStyle)), ds, ws, 1))
//		return
//	}
//	if *align == "center" {
//
//		ws := GetWinSize()
//		ds := GetArtWidth(input, GetCharsWidth(PrepareBan(bannerStyle)))
//		fmt.Println(MakeArtAligned(input, GetChars(PrepareBan(bannerStyle)), ds, ws, 2))
//		return
//	}
//	if *align == "justify" {
//
//		ws := GetWinSize()
//		ds := GetArtWidth(input, GetCharsWidth(PrepareBan(bannerStyle)))
//		fmt.Println(MakeArtJustified(input, GetChars(PrepareBan(bannerStyle)), ds, ws))
//		return
//	}
//
//	if *reverse != "default" {
//
//		file, err := os.Open(*reverse)
//		if err != nil {
//			fmt.Println("Error opening the file:", err)
//			return // Exit the program on error
//		}
//		defer pkg(file *os.File) {
//			err := file.Close()
//			if err != nil {
//				fmt.Println("Error closing the file:", err)
//			}
//		}(file)
//		source := FileToVariable(file)
//		emptyCols := RemoveValidSPaceIndex(GetEmptyCols(source))
//		charMap := GetInputChars(ArtToSingleLine(source), emptyCols)
//		mapStandard := GetChars(PrepareBan("standard"))
//		mapShadow := GetChars(PrepareBan("shadow"))
//		mapThinkertoy := GetChars(PrepareBan("thinkertoy"))
//		AsciiToChars(charMap, mapStandard, mapShadow, mapThinkertoy)
//	}
//	// test is for testing and debugging
//	if *test {
//		fmt.Println("Reserved for testing and debugging.")
//	 else {
// default output
//
//		fmt.Println(MakeArt(input, GetChars(PrepareBan(bannerStyle))))
