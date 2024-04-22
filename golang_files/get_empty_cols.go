package asciiartweb

import "fmt"

func GetEmptyCols(source []string) []int {
	fmt.Println("entering GetEmptyCols")
	fmt.Println("------------------------------------------------")

	source = ArtToSingleLine(source)

	fmt.Println("source:", source)

	var emptyCols []int

	//loop over each line's character
	for i := 0; i < len(source[0]); i++ {
		empty := true
		//then loop over the lines of the source ascii art text
		for j := 0; j < len(source); j++ {
			if source[j][i] != ' ' {
				empty = false
			}
		}
		if empty {
			emptyCols = append(emptyCols, i)
		}
	}
	return emptyCols
}

// Remove indices for valid spaces, before the end space
func RemoveValidSpaceIndex(indices []int) []int {
	fmt.Println("entering RemoveValidSpaceIndex")
	fmt.Println("------------------------------------------------")
	for i := 0; i < len(indices)-1; i++ {
		if len(indices)-i > 6 {
			if indices[i] == (indices[i+6])-6 {
				indices = append(indices[:i+1], indices[i+6:]...)
			}
		}
	}
	return indices
}
