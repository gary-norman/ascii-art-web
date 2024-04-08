package ascii_art_web

import (
	"fmt"
)

/*
	Reverse Plan:
	1. loop over text file:
		-loop over each character in a line
			-how to distinguish where each character starts and ends?
				shown below
		-loop over the amount of lines
	2. when looping, check if each ascii character's string matches the ascii map
	3. if all 8 checks for each line are equal to the same letter, store that letter
	4. create a new string and keep adding each letter to create the word to print
	5. print the word

	How to separate the characters in the ascii art
	1. create a map for each character style - standard, shadow, thinkertoy
	2. use empty space across all 8 lines as a separator
	3. extract these as an array of indexes
	4. add a star at these indexes in the lines
	5. strings.Split using * as a separator
	6. go over the arrays and add the characters for each letter/symbol together and map them
	7. compare the input text ascii map with each of the prepared maps
	8. Print each word
*/

// Compares getChar and getInputChar and prints the string to the terminal
func AsciiToChars(input, standard, shadow, thinkertoy map[int][]string) {
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
	for i := 0; i < len(output); i++ {
		fmt.Printf("%v", string(rune(output[i][0])))
	}
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
