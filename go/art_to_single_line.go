package ascii_art_web

import "fmt"

// Places each line of characters from FileToVariable on a single line, delineated by "** "
func ArtToSingleLine(source []string) []string {
	fmt.Println("entering ArtToSingleLine")
	fmt.Println("------------------------------------------------")
	var output []string

	for index, char := range source {
		fmt.Println(index, char)
	}

	fmt.Println("len of source:", len(source))
	//if source has only 8 lines, it means the words are in one line, not separated by \n
	if len(source) == 9 {
		//then return it back
		source = source[:8]
	}

	//loop over each line and asign source's lines to output
	for i := 0; i < 8; i++ {
		output = append(output, source[i])
	}

	//if output is longer than 8, means that words are separated by \n
	if len(source) > 9 {
		// make source everything after the 8
		source = source[10:]
	}

	//loop until source has no lines
	x := 0
	for len(source) > 0 {
		for i := 0; i < 8; i++ {
			//append every 8th line with * and # for figuring out the placement of end of line (between characters) and end of line for each part of the ascii character (out of 8 lines)
			output = append(output, output[i+x]+"* "+"# "+source[i])
			// fmt.Println("output:", output)
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
	// fmt.Println("output:", output)
	return output
}
