package ascii_art_web

import (
	"flag"
	"os"
)

func PrepareBanner(style string) []string {
	args := flag.Args()
	lenOfArg := len(flag.Args())
	var scanned []string
	var newStyle string

	//if empty string is passed, most functions work, except for reverse
	if style == "" {

		//in case of color flag, with 4 arguments: file directory,"--color=<color>", "toColour" and "fullString"
		if len(os.Args) == 4 {
			newStyle = "standard"

			//in case of 2 arguments: fullString and style
		} else if lenOfArg > 1 {
			if args[1] == "thinkertoy" {
				newStyle = "thinkertoy"
			} else if args[1] == "standard" {
				newStyle = "standard"
			} else if args[1] == "shadow" {
				newStyle = "shadow"
			}
		} else {
			//otherwise set to standard, so the other functions use a standard font
			newStyle = "standard"
		}
		scanned = PrepareFile(newStyle)

		//in case of reverse, style is passed in the string and set accordingly to scanned
	} else {
		newStyle = style
		scanned = PrepareFile(newStyle)
	}
	return scanned
}
