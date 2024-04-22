package asciiartweb

func AsciiMap(lines []string) map[int][]string {

	var asciiMap = map[int][]string{}
	id := 31
	for _, line := range lines {
		// line = strings.TrimSpace(line)
		if line == "" {
			id++
		} else {
			asciiMap[id] = append(asciiMap[id], line)
		}
	}
	return asciiMap
}

// Map the ascii characters provided in the reverse flag, zero indexed
func CharMap(source []string, indices []int) map[int][]string {
	charMap := make(map[int][]string)
	startIndex := 0
	for id := range indices {
		for _, line := range source {
			charMap[id] = append(charMap[id], line[startIndex:indices[id]]+" ")
		}
		startIndex = indices[id] + 1
	}
	return charMap
}
