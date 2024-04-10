package main

import (
	ascii_art_web "ascii_art_web/go"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func homePage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		return
	}
}

func handleRequests() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("icons"))))
	http.Handle("/ascii_styles/", http.StripPrefix("/ascii_styles/", http.FileServer(http.Dir("ascii_styles"))))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/process", processor)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var artInput, artOutput, outputResult string
	//standardMap := ascii_art_web.AsciiMap(ascii_art_web.PrepareBanner("standard"))
	//shadowMap := ascii_art_web.AsciiMap(ascii_art_web.PrepareBanner("shadow"))
	//thinkertoyMap := ascii_art_web.AsciiMap(ascii_art_web.PrepareBanner("thinkertoy"))

	chosenInput := r.FormValue("generate")
	chosenStyle := r.FormValue("banner")
	chosenColor := r.FormValue("colors")
	colorInput := r.FormValue("colour-text")
	chosenAlign := r.FormValue("text-align")
	defaultValue := "default"

	if ascii_art_web.IsFilePresent(w, r) {
		emptyCols := ascii_art_web.RemoveValidSpaceIndex(ascii_art_web.GetEmptyCols(ascii_art_web.ArtFromFileLines(w, r)))
		fmt.Println("emptyCols", emptyCols)

		charMap := ascii_art_web.CharMap(ascii_art_web.ArtToSingleLine(ascii_art_web.ArtFromFileLines(w, r)), emptyCols)
		fmt.Println("charMap:", charMap)

		artInput = ascii_art_web.ArtFromFile(w, r)
		fmt.Println("conditi n: artInput is:		", artInput)
		artOutput = ascii_art_web.CheckReverse(w, r)
		fmt.Println("condition: artOutput is:		", artOutput)
	}
	if colorInput == "" && chosenInput != "" {

		colorInput = chosenInput
	}
	fmt.Println("chosenInput is:		", chosenInput)
	fmt.Println("chosenStyle is:		", chosenStyle)
	fmt.Println("chosenColor is:		", chosenColor)
	fmt.Println("colorInput is:		", colorInput)
	fmt.Println("artInput is:		", artInput)
	fmt.Println("artOutput is:		", artOutput)
	fmt.Println("chosenAlign is:		", chosenAlign)
	fmt.Println("------------------------------------------------")

	if artInput != "" {
		outputResult = artInput
		fmt.Println("file present... outputResult is now = artInput -> outputResult:", outputResult)

	} else {
		outputResult = ascii_art_web.RunAscii(chosenInput, chosenStyle, chosenColor, colorInput, defaultValue, chosenAlign, artInput)
	}

	fmt.Println("------------------------------------------------------------------------------------------------")
	d := struct {
		InputText  string
		Style      string
		Color      string
		ColorWord  string
		FileWant   string
		InputAlign string
		ArtToText  string
		TextToArt  string
	}{
		InputText:  chosenInput,
		Style:      chosenStyle,
		Color:      chosenColor,
		FileWant:   defaultValue,
		ColorWord:  colorInput,
		InputAlign: chosenAlign,
		ArtToText:  artOutput,
		TextToArt:  outputResult,
	}
	tpl.ExecuteTemplate(w, "result.html", d)
}
