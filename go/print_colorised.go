package ascii_art_web

import (
	"fmt"
	"html"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Orange = "\033[38;5;208m"
var Grey = "\033[37m"

func PrintColorised(colour string, asciiStr string) string {
	var displayColor, printColor string
	newColoredString := ""
	switch colour {
	case "red":
		displayColor = "color: red;"
		printColor = Red
	case "#ff0000":
		displayColor = "color: red;"
		printColor = Red
	case "rgb(255, 0, 0)":
		displayColor = "color: red;"
		printColor = Red
	case "hsl(0, 100%, 50%)":
		displayColor = "color: red;"
		printColor = Red
	case "blue":
		displayColor = "color: blue;"
		printColor = Blue
	case "green":
		displayColor = "color: green;"
		printColor = Green
	case "yellow":
		displayColor = "color: yellow;"
		printColor = Yellow
	case "purple":
		displayColor = "color: purple;"
		printColor = Purple
	case "cyan":
		displayColor = "color: cyan;"
		printColor = Cyan
	case "grey":
		displayColor = "color: grey;"
		printColor = Grey
	case "orange":
		displayColor = "color: orange;"
		printColor = Orange
	default:
		displayColor = ""
		printColor = Reset
	}
	if displayColor == "" {
		newColoredString = fmt.Sprintf(`<span style="%s">%s</span>`, displayColor, asciiStr)

	} else {
		newColoredString = asciiStr
	}

	printString := printColor + asciiStr + Reset
	fmt.Print(printString)

	// HTML escape characters
	newColoredString = html.UnescapeString(newColoredString)
	//fmt.Println("newColoredString is:", newColoredString)
	return newColoredString

}

//func ColoriseLetter(asciiMap map[int][]string, colour, toColour, word string) string {
//	var displayColor, printColor string
//	newColoredString := ""
//	switch colour {
//	case "red":
//		displayColor = "color: red;"
//		printColor = Red
//	case "#ff0000":
//		displayColor = "color: red;"
//		printColor = Red
//	case "rgb(255, 0, 0)":
//		displayColor = "color: red;"
//		printColor = Red
//	case "hsl(0, 100%, 50%)":
//		displayColor = "color: red;"
//		printColor = Red
//	case "blue":
//		displayColor = "color: blue;"
//		printColor = Blue
//	case "green":
//		displayColor = "color: green;"
//		printColor = Green
//	case "yellow":
//		displayColor = "color: yellow;"
//		printColor = Yellow
//	case "purple":
//		displayColor = "color: purple;"
//		printColor = Purple
//	case "cyan":
//		displayColor = "color: cyan;"
//		printColor = Cyan
//	case "grey":
//		displayColor = "color: grey;"
//		printColor = Grey
//	case "orange":
//		displayColor = "color: orange;"
//		printColor = Orange
//	default:
//		displayColor = ""
//		printColor = Reset
//	}
//	if displayColor == "" {
//		newColoredString = fmt.Sprintf(`<span style="%s">%s</span>`, displayColor, asciiStr)
//
//	} else {
//		newColoredString = word
//	}
//
//	var coloredArray []string
//
//	for indexA, charA := range  {
//		for indexB, charB := range word {
//			if ascii, ok := asciiMap[int(word[char])]; ok {
//
//			}
//		}
//	}
//
//
//	printString := printColor + word + Reset
//	fmt.Print(printString)
//
//	// HTML escape characters
//	newColoredString = html.UnescapeString(newColoredString)
//	//fmt.Println("newColoredString is:", newColoredString)
//	return newColoredString
//
//}
