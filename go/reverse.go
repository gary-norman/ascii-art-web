package ascii_art_web

import (
	"fmt"
)

// AsciiToChars Compares getChar and getInputChar and prints the string to the terminal
func AsciiToChars(input, standard, shadow, thinkertoy map[int][]string) {
	fmt.Println("entering AsciiToChars")
	fmt.Println("------------------------------------------------")

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
		fmt.Printf("output[i][0]: %v\n", string(rune(output[i][0])))
	}
	fmt.Println("reverse.go -> output is:", output)
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
