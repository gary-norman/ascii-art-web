package ascii_art_web

import (
	"fmt"
	"strings"
)

//

func PrintAligned(align string, sep []string, spaces int, s string) {
	// Calculate window size to get terminal's width
	winSize := GetWinSize()
	winWidth := int(winSize.Col)

	//Get string's length
	sLen := len(s)

	// Calculate total padding required to center the string
	centerPadding := (winWidth-sLen)/2 - 1
	rightPadding := winWidth - sLen - 1
	var justifyPadding int
	if len(sep) > 1 {
		justifyPadding = (winWidth - sLen) / (len(sep) - 1)
	}

	//handle each align case
	switch align {
	case "right":
		fmt.Println(strings.Repeat(" ", rightPadding), s)
	case "center":
		fmt.Println(strings.Repeat(" ", centerPadding), s)
	case "justify":
		var line string

		//loop over the string line and add correct padding after each word
		if len(sep) == 1 {
			fmt.Println(strings.Repeat(" ", centerPadding), s)
		} else {
			for _, word := range sep {
				line = line + word + strings.Repeat(" ", justifyPadding)
			}
			//trim the spaces after the last word
			line = strings.TrimRight(line, " ")
			fmt.Println(line)
		}
	default:
		fmt.Println(s)
	}
}
