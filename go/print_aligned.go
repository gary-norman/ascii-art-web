package ascii_art_web

//

//func AlignWords(charMap map[int][]string, words []string, align string) {
//	var stringArray []string
//	var wordCounter = 0
//	for i := 0; i < len(charMap); i++ {
//
//		if i == 0 || charMap[i][] == 32 {
//
//		}
//
//	}
//}
// commented this all so I could get mine to run

func AlignJustified(align string, sep []string, s string) string {
	line := "<pre class=\"justify\">"

	//loop over the string line and add correct padding after each word
	for _, word := range sep {
		line = line + word + "</pre>"
	}
	return line
}
