package api

import "net/http"

func HandleRequestsGary() {
	http.HandleFunc("/", HomePageGary)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("icons"))))
	//http.Handle("/pkg/", http.StripPrefix("/pkg/", http.FileServer(http.Dir("pkg"))))
	http.Handle("/ascii_styles/", http.StripPrefix("/ascii_styles/", http.FileServer(http.Dir("ascii_styles"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
