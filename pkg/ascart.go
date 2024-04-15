package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// Winsize is a struct that stores the height and width of the terminal.
type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16 // unused
	Ypixel uint16 // unused
}

// *************** Global Functions *************** //

// GetWinSize populates the Winsize structure
func GetWinSize() Winsize {
	// Get the file descriptor for stdout
	fd := syscall.Stdout

	// Create an instance of Winsize
	var ws Winsize

	// Use the TIOCGWINSZ ioctl system call to get the window size
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&ws)))
	if err != 0 {
		fmt.Println("Error getting terminal size:", err)
	}
	return ws
}

// PrepareBan import standard.txt as the default ascii style, with ability to change it using 2nd argument *
func PrepareBan(bannerStyle string) []string {
	if bannerStyle == "" {
		bannerStyle = "standard"
	}
	style := bannerStyle
	file, err := os.Open("ascii_styles/" + style + ".txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing the file:", err)
		}
	}(file)
	scanned := bufio.NewScanner(file)
	scanned.Split(bufio.ScanLines)
	var source []string
	for scanned.Scan() {
		source = append(source, scanned.Text())
	}
	return source
}

// FileToVariable takes a file as input and returns it as a slice of strings
func FileToVariable(file *os.File) []string {
	scanned := bufio.NewScanner(file)
	scanned.Split(bufio.ScanLines)
	var source []string
	for scanned.Scan() {
		source = append(source, scanned.Text())
	}
	return source
}

// CompareSlices compares two slices for equality.
func CompareSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false // Slices are of different lengths
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false // Elements at the same position are different
		}
	}
	return true // Slices are equal
}

// Contains checks if a slice contains a specific element.
func Contains(slice []rune, item rune) bool {
	for _, v := range slice {
		if v == item {
			return true // Found the item
		}
	}
	return false // Item not found
}

// CheckFlagSkipsEquals checks if the value passed to the flag uses <flag>=<argument> format
func CheckFlagSkipsEquals(flag string) bool {
	usedEqualsSyntax := true
	for _, arg := range os.Args[1:] {
		// Check if the argument starts with the flag name followed by "="
		if strings.HasPrefix(arg, flag) {
			usedEqualsSyntax = false
		}
	}
	return usedEqualsSyntax
}

// *************** Helper Functions *************** //

// ArtToSingleLine Places each line of characters from FileToVariable on a single line, delineated by "** "
func ArtToSingleLine(source []string) []string {
	var output []string
	if len(source) == 8 {
		return source
	}
	for i := 0; i < 8; i++ {
		output = append(output, source[i])
	}
	if len(source) > 8 {
		source = source[8:]
	}
	x := 0
	for len(source) > 0 {
		for i := 0; i < 8; i++ {
			output = append(output, output[i+x]+"* "+"# "+source[i])
		}
		x += 8
		if len(source) > 8 {
			source = source[8:]
		} else {
			source = nil
		}
	}
	for len(output) > 8 {
		output = output[8:]
	}
	return output
}

// GetEmptyCols Get the index of the final space of each character in the reverse flag
func GetEmptyCols(source []string) []int {
	source = ArtToSingleLine(source)
	var emptyCols []int
	for i := 0; i < len(source[0]); i++ {
		empty := true
		for j := 0; j < len(source); j++ {
			if source[j][i] != 32 {
				empty = false
			}
		}
		if empty == true {
			emptyCols = append(emptyCols, i)
		}
	}
	return emptyCols
}

// RemoveValidSPaceIndex Remove indices for valid spaces, before the end space. Valid spaces are 6 characters wide.
func RemoveValidSPaceIndex(indices []int) []int {
	for i := 0; i < len(indices)-1; i++ {
		if len(indices)-i > 6 {
			if indices[i] == (indices[i+6])-6 {
				indices = append(indices[:i+1], indices[i+6:]...)
			}
		}
	}
	return indices
}

// GetChars Map the ascii characters provided in the style.txt file, indexed by ascii code
func GetChars(source []string) map[int][]string {
	charMap := make(map[int][]string)
	id := 31
	for _, line := range source {
		if string(line) == "" {
			id++
		} else {
			charMap[id] = append(charMap[id], line)
		}
	}
	return charMap
}

// GetInputChars Map the ascii characters provided in the reverse flag, zero indexed
func GetInputChars(source []string, indices []int) map[int][]string {
	charMap := make(map[int][]string)
	startIndex := 0
	for id := range indices {
		for _, line := range source {
			charMap[id] = append(charMap[id], line[startIndex:indices[id]]+" ")
		}
		startIndex = indices[id] + 1
	}
	return charMap
}

// AsciiToChars Compares getChar and getInputChar and prints the string to the terminal
func AsciiToChars(input, standard, shadow, thinkertoy map[int][]string) string {
	output := make(map[int][]int)
	var newLine1 []string
	for i := 0; i < 8; i++ {
		newLine1 = append(newLine1, "* ")
	}
	var newLine2 []string
	for i := 0; i < 8; i++ {
		newLine2 = append(newLine2, "# ")
	}
	slash := 92
	n := 110
	styles := []map[int][]string{standard, shadow, thinkertoy}
	for _, style := range styles {
		for key1, slice1 := range input {
			for key2, slice2 := range style {
				if CompareSlices(slice1, newLine1) {
					output[key1] = append(output[key1], slash)
				}
				if CompareSlices(slice1, newLine2) {
					output[key1] = append(output[key1], n)
				}
				if CompareSlices(slice1, slice2) {
					output[key1] = append(output[key1], key2)
				}
			}
		}
	}
	var outString string
	//for _, arr := range output {
	//	for _, digit := range arr {
	//		char := rune(digit)
	//		outString += string(char)
	//	}
	//}
	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[i]); j++ {
			char := rune(output[i][j])
			outString += string(char)
		}
	}
	return outString
}

// GetCharsWidth Determine the width of each individual ascii art character
func GetCharsWidth(source []string) map[int]int {
	charWidthMap := make(map[int]int)
	id := 31
	for _, line := range source {
		if string(line) == "" {
			id++
		} else {
			charWidthMap[id] = len(line)
		}
	}
	return charWidthMap
}

// GetArtWidth Determine the width of each line that gets printed to the terminal (without EOL)
func GetArtWidth(origString string, y map[int]int) []int {
	var width []int
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	for i := 0; i < len(wordSlice); i++ {
		sum := 0
		for _, char := range wordSlice[i] {
			sum += y[int(char)]
			//for _, num := range y[int(char)] {
			//	sum += num
			//}
		}
		width = append(width, sum)
	}
	return width
}

// *************** Final Output Functions *************** //

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
func RunMakeArt() {
	MakeArt("test", GetChars(PrepareBan("standard")))
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

// MakeArtJustified Transform the input text origString to the output art, line by line, with justified content
func MakeArtJustified(origString string, y map[int][]string) string {
	var art string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	var words []string
	for i := 0; i < len(wordSlice); i++ {
		words = append(words, strings.Split(wordSlice[i], " ")...)
	}
	//fmt.Printf("Words = %v\n", words)
	//fmt.Printf("Wordslength = %v\n", len(words))

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

//pkg main() {
//	// ? flag definitions
//	reverse := flag.String("reverse", "default",
//		"Convert ascii art from a specified file into a string of characters.")
//	skipsFlagReverse := CheckFlagSkipsEquals("--reverse=")
//	color := flag.String("color", "default",
//		"Format the output into a specified colour, either the entire text, or limited to specified characters.")
//	skipsFlagColor := CheckFlagSkipsEquals("--color=")
//	output := flag.String("output", "default",
//		"Save the output to the specified filename")
//	skipsFlagOutput := CheckFlagSkipsEquals("--output=")
//	align := flag.String("align", "default",
//		"Align the output to a specified alignment.")
//	skipsFlagAlign := CheckFlagSkipsEquals("--align=")
//	help := flag.Bool("help", false, "Provide the user with a help file.")
//	test := flag.Bool("test", false, "testing")
//	flag.Parse() // parse the flags so that they can be used
//	// define args to be non-flag arguments
//	additionalArgs := flag.Args() // tell the program to treat every argument following the flag as arg[o...]
//	var input string
//	if len(additionalArgs) > 0 {
//		input = additionalArgs[0]
//	}
//	var bannerStyle string
//	if len(additionalArgs) == 2 {
//		bannerStyle = additionalArgs[1]
//	}
//	// call the functions depending on the flag
//	if *help {
//		cmd := exec.Command("cat", "help.txt")
//		cmd.Stdout = os.Stdout
//		if err := cmd.Run(); err != nil {
//			fmt.Println("could not run command: ", err)
//		}
//	}
//	if *color != "default" {
//		var colored string
//		var colorAll bool
//		var colSLice []rune
//		colorAll = true
//		if skipsFlagColor {
//			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
//			return
//		}
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
//		if skipsFlagOutput {
//			fmt.Println("Usage: go run . [OPTION] [STRING] [STYLE]\n\nEX: go run . --output=<filename> \"something\" standard")
//			return
//		}
//		err := os.WriteFile(*output, []byte(MakeArt(input, GetChars(PrepareBan(bannerStyle)))+"\n"), 0644)
//		if err != nil {
//			fmt.Println("Error writing to the file:", err)
//			return // Exit the program on error
//		}
//		fmt.Println(MakeArt(input, GetChars(PrepareBan(bannerStyle))))
//		fmt.Printf("Output has been saved to %v\n", *output)
//		return
//	}
//	if *align == "left" {
//		if skipsFlagAlign {
//			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
//			return
//		}
//		fmt.Println(MakeArt(input, GetChars(PrepareBan(bannerStyle))))
//		return
//	}
//	if *align == "right" {
//		if skipsFlagAlign {
//			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
//			return
//		}
//		ws := GetWinSize()
//		ds := GetArtWidth(input, GetCharsWidth(PrepareBan(bannerStyle)))
//		fmt.Println(MakeArtAligned(input, GetChars(PrepareBan(bannerStyle)), ds, ws, 1))
//		return
//	}
//	if *align == "center" {
//		if skipsFlagAlign {
//			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
//			return
//		}
//		ws := GetWinSize()
//		ds := GetArtWidth(input, GetCharsWidth(PrepareBan(bannerStyle)))
//		fmt.Println(MakeArtAligned(input, GetChars(PrepareBan(bannerStyle)), ds, ws, 2))
//		return
//	}
//	if *align == "justify" {
//		if skipsFlagAlign {
//			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
//			return
//		}
//		ws := GetWinSize()
//		ds := GetArtWidth(input, GetCharsWidth(PrepareBan(bannerStyle)))
//		fmt.Println(MakeArtJustified(input, GetChars(PrepareBan(bannerStyle)), ds, ws))
//		return
//	}
//	if *reverse != "default" {
//		if skipsFlagReverse {
//			fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
//			return
//		}
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
//		// default output
//		if len(additionalArgs) > 2 {
//			fmt.Println("Usage: go run . [STRING] [STYLE] (optional)\n\nEX: go run . \"something\" thinkertoy.\nAvailable styles are standard, shadow, and thinkertoy.")
//			return
//		}
//		fmt.Println(MakeArt(input, GetChars(PrepareBan(bannerStyle))))
//	}
//}
