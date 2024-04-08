package ascii_art_web

func PrepareBanner(style string) []string {
	var scanned []string

	//if empty string is passed, most functions work, except for reverse
	if style == "" {

		scanned = PrepareFile("standard")

		//in case of reverse, style is passed in the string and set accordingly to scanned
	} else {
		scanned = PrepareFile(style)
	}
	return scanned
}
