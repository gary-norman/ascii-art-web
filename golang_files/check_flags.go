package asciiartweb

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func CheckFlagsAndArgs(colour string, output string, reverse string, align string) bool {
	//process default number of arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: golang_files run . [STRING] [BANNER]\n\nExample: golang_files run . something standard")
		return false
	}
	tag := os.Args[1]
	otherArgs := flag.Args()
	lenOfArg := len(otherArgs)

	//if the flag (tag) doesn't include "=", return false and print correct message
	if !strings.Contains(tag, "=") {
		//check for correct flags
		if reverse != "default" {
			fmt.Println("Usage: golang_files run . [OPTION] [STRING] [BANNER]\n\nExample: golang_files run . --output=<fileName.txt> something standard")
			return false
		} else if output != "default" {
			fmt.Println("Usage: golang_files run . [OPTION] [STRING] [BANNER]\n\nExample: golang_files run . --reverse=<fileName>")
			return false
		} else if colour != "default" {
			fmt.Println("Usage: golang_files run . [OPTION] [STRING] [BANNER]\n\nExample: golang_files run . --color=<color> <letters to be colored> \"something\"")
			return false
		} else if align != "default" {
			fmt.Println("Usage: golang_files run . [OPTION] [STRING] [BANNER]\n\nExample: golang_files run . --align=right something standard")
			return false
		}
	}

	//after passing all flags correctly, if args are more than 2, return false and print message
	if lenOfArg > 2 {
		fmt.Println("Usage: golang_files run . [STRING] [BANNER]\n\nExample: golang_files run . something standard")
		return false
	}
	return true
}
