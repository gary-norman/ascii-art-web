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

	var inputArt, artToText, textToArt string
	var charMap map[int][]string
	//standardMap := ascii_art_web.AsciiMap(ascii_art_web.PrepareBanner("standard"))
	//shadowMap := ascii_art_web.AsciiMap(ascii_art_web.PrepareBanner("shadow"))
	//thinkertoyMap := ascii_art_web.AsciiMap(ascii_art_web.PrepareBanner("thinkertoy"))

	inputText := r.FormValue("generate")
	chosenStyle := r.FormValue("banner")
	chosenColor := r.FormValue("colors")
	colorWord := r.FormValue("colour-text")
	chosenAlign := r.FormValue("text-align")
	defaultValue := "default"
	yourArt := "Your art says: "

	if ascii_art_web.IsFilePresent(w, r) {
		emptyCols := ascii_art_web.RemoveValidSpaceIndex(ascii_art_web.GetEmptyCols(ascii_art_web.ArtFromFileLines(w, r)))
		fmt.Println("emptyCols", emptyCols)

		charMap = ascii_art_web.CharMap(ascii_art_web.ArtToSingleLine(ascii_art_web.ArtFromFileLines(w, r)), emptyCols)
		fmt.Println("charMap:", charMap)

		inputArt = ascii_art_web.ArtFromFile(w, r)
		fmt.Println("conditi n: inputArt is:		", inputArt)
		artToText = ascii_art_web.CheckReverse(w, r)
		yourArt += artToText
		fmt.Println("condition: artToText is:		", artToText)
	}

	fmt.Println("inputText is:		", inputText)
	fmt.Println("chosenStyle is:		", chosenStyle)
	fmt.Println("chosenColor is:		", chosenColor)
	fmt.Println("colorWord is:		", colorWord)
	fmt.Println("inputArt is:		", inputArt)
	fmt.Println("artToText is:		", artToText)
	fmt.Println("chosenAlign is:		", chosenAlign)
	fmt.Println("------------------------------------------------")

	if inputArt != "" {
		textToArt = inputArt
		fmt.Println("file present... textToArt is now = inputArt -> textToArt:", textToArt)

	} else {
		textToArt = ascii_art_web.RunAscii(inputText, chosenStyle, chosenColor, colorWord, defaultValue, chosenAlign, inputArt, charMap)
	}

	if colorWord == "" && inputText != "" {
		colorWord = inputText
		textToArt = "<pre class=\"" + chosenColor + "\">" + textToArt + "</pre>"
	} else {
		textToArt = "<pre class=\"\">" + textToArt + "</pre>"
	}

	fmt.Println("------------------------------------------------------------------------------------------------")
	d := struct {
		InputText   string //"hello"
		ChosenStyle string //"standard"
		ChosenColor string //"red"
		ColorWord   string //"lo"
		ChosenAlign string //left, right etc
		YourArt     string //"Your Art says: "blue lagoon"
		ArtToText   string //"blue lagoon"
		TextToArt   string //Your art: "____||||_|_|__|_|_|__|_"
	}{
		InputText:   inputText,
		ChosenStyle: chosenStyle,
		ChosenColor: chosenColor,
		ColorWord:   colorWord,
		ChosenAlign: chosenAlign,
		YourArt:     yourArt,
		ArtToText:   artToText,
		//TextToArt:  "<pre>this</pre><pre>is</pre><pre>a</pre><pre>test</pre>",
		TextToArt: textToArt,
	}
	tpl.ExecuteTemplate(w, "result.html", d)
}
