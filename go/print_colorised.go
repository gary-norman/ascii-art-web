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
var Gray = "\033[37m"
var White = "\033[97m"

func PrintColorised(colour string, asciiStr string) string {
	var extraColour string
	newColoredString := ""
	switch colour {
	case "red":
		extraColour = "color: red;"
	case "#ff0000":
		extraColour = "color: red;"
	case "rgb(255, 0, 0)":
		extraColour = "color: red;"
	case "hsl(0, 100%, 50%)":
		extraColour = "color: red;"
	case "blue":
		extraColour = "color: blue;"
	case "green":
		extraColour = "color: green;"
	case "yellow":
		extraColour = "color: yellow;"
	case "purple":
		extraColour = "color: purple;"
	case "cyan":
		extraColour = "color: cyan;"
	case "grey":
		extraColour = "color: grey;"
	case "orange":
		extraColour = "color: orange;"
	default:
		extraColour = ""
	}

	newColoredString = fmt.Sprintf(`<span style="%s">%s</span>`, extraColour, asciiStr)

	// HTML escape characters
	newColoredString = html.UnescapeString(newColoredString)
	fmt.Println("newColoredString is:", newColoredString)
	return newColoredString

}
