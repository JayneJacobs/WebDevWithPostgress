package main

import (
	"net/http"
)

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "public"+r.URL.Path)
// 	})
// 	http.ListenAndServe(":8000", nil)
// }
// The following replaces above

func main() {
	http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}
