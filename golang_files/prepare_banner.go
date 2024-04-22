package asciiartweb

func PrepareBanner(style string) []string {
	var scanned []string

	//if empty string is passed, most functions work, except for reverse
	if style == "" {
		scanned = PrepareFile("standard")
	} else {
		scanned = PrepareFile(style)
	}
	return scanned
}
