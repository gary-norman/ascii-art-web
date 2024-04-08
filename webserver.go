package main

import "ascii_art_web/api"

//pkg methodHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "GET":
//		fmt.Fprintf(w, "GET request received")
//	case "POST":
//		fmt.Fprintf(w, "POST request received")
//	default:
//		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
//	}
//}

func main() {
	api.HandleRequestsGary()

}
