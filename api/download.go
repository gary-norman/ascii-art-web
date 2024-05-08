package api

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"strconv"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entering Downloader func")
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	file, err := os.Open("arttofile/yourart.txt")
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	fileInfo, err := os.Stat("arttofile/yourart.txt")
	if err != nil {
		http.Error(w, "Error Stating file", http.StatusInternalServerError)
		return
	}
	cd := mime.FormatMediaType("attachment", map[string]string{"filename": fileInfo.Name()})
	w.Header().Set("Content-Disposition", cd)
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}
