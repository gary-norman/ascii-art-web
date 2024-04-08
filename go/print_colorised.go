package ascii_art_web

import "fmt"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Orange = "\033[38;5;208m"
var Gray = "\033[37m"
var White = "\033[97m"

func PrintColorised(colour string, asciiStr string) {
	var extraColour string
	switch colour {
	case "red":
		extraColour = Red
	case "#ff0000":
		extraColour = Red
	case "rgb(255, 0, 0)":
		extraColour = Red
	case "hsl(0, 100%, 50%)":
		extraColour = Red
	case "blue":
		extraColour = Blue
	case "green":
		extraColour = Green
	case "yellow":
		extraColour = Yellow
	case "purple":
		extraColour = Purple
	case "cyan":
		extraColour = Cyan
	case "gray":
		extraColour = Gray
	case "orange":
		extraColour = Orange
	default:
		extraColour = Reset
	}

	newString := extraColour + asciiStr + Reset
	fmt.Print(newString)
}
