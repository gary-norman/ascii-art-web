package pkg

import (
	"fmt"
	"net/http"
)

func Download(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering download")

	w.Header().Add("Content-Disposition", "attachment;filename=\"yourart.txt")
	http.ServeFile(w, r, "arttofile/yourart.txt")
}
