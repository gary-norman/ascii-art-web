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

func wordsToPre(words []string, y map[int][]string, wordSlice []string) string {
	var art string
	if len(words) > 1 {
		var line string
		for _, word := range words {
			line = "<pre>"
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
			//fmt.Printf("Added </pre>\n")
			art += line
			//fmt.Printf("%v += %v\n", art, line)
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

// MakeArtJustified Transform the input text origString to the output art, line by line, with justified content
func MakeArtJustified(origString string, y map[int][]string, align string) (string, string) {
	var art string
	var justification string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	var lines [][]string
	for _, str := range wordSlice {
		splitStr := strings.Split(str, "\n")
		lines = append(lines, splitStr)
	}
	if len(lines) > 1 {
		for _, newLine := range lines {
			var newWords []string
			for _, word := range newLine {
				newWords = append(newWords, strings.Split(word, " ")...)
			}
			art += "<div class=\"justifiedOutput" + align + "\">"
			art += wordsToPre(newWords, y, newLine)
			art += "</div>"
			justification = "Multiline"
		}
	} else {
		for _, newLine := range lines {
			var newWords []string
			for _, word := range newLine {
				newWords = append(newWords, strings.Split(word, " ")...)
			}
			art += wordsToPre(newWords, y, newLine)
			justification = align
		}
	}
	return art, justification
}

// MakeArtColorized Transform the input text origString to the output art, line by line, colorizing specified text
func MakeArtColorized(origString string, y map[int][]string, letters []rune, color string, colorAll bool) string {
	//var specifiedColor string
	//reset := "\033[0m"
	//switch color {
	//case "red":
	//	specifiedColor = "\033[31m"
	//case "#ff0000":
	//	specifiedColor = "\033[31m"
	//case "rgb(255, 0, 0)":
	//	specifiedColor = "\033[31m"
	//case "hsl(0, 100%, 50%)":
	//	specifiedColor = "\033[31m"
	//case "green":
	//	specifiedColor = "\033[32m"
	//case "#00ff00":
	//	specifiedColor = "\033[32m"
	//case "rgb(0, 255, 0)":
	//	specifiedColor = "\033[32m"
	//case "hsl(120, 100%, 50%)":
	//	specifiedColor = "\033[32m"
	//case "yellow":
	//	specifiedColor = "\033[33m"
	//case "#f0ff00":
	//	specifiedColor = "\033[33m"
	//case "rgb(240, 255, 0)":
	//	specifiedColor = "\033[33m"
	//case "hsl(64, 100%, 50%)":
	//	specifiedColor = "\033[33m"
	//case "blue":
	//	specifiedColor = "\033[34m"
	//case "#0000ff":
	//	specifiedColor = "\033[34m"
	//case "rgb(0, 0, 255)":
	//	specifiedColor = "\033[34m"
	//case "hsl(240, 100%, 50%)":
	//	specifiedColor = "\033[34m"
	//case "orange":
	//	specifiedColor = "\033[38;5;208m"
	//case "#f9690e":
	//	specifiedColor = "\033[38;5;208m"
	//case "rgb(249, 105, 14)":
	//	specifiedColor = "\033[38;5;208m"
	//case "hsl(23, 100%, 50%)":
	//	specifiedColor = "\033[38;5;208m"
	//default:
	//	fmt.Print("\nAvailable colors are " + "\033[31m" + "red" + reset + ", " +
	//		"\033[32m" + "green" + reset + "," + "\033[33m" + "yellow" + reset + ", " +
	//		"\033[38;5;208m" + "orange" + reset + ", and " + "\033[34m" + "blue" + reset + ".\n")
	//}
	var art string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	var line string
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
