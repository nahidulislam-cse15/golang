package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Println(w.Header().Get("Content-Type"))
	fmt.Fprint(w, "<h1>Welcome to Master Academy</h1>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, r.URL.Path)
}

// func about(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type:", "text/plain")
// 	fmt.Fprint(w, "<h1>Welcome to Master Academy about page</h1>")
// }
// func contact(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<h1>Welcome to Master Academy contact page</h1>")
// }
// func blog(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<h1>Wlcome to Master Academy blog page</h1>")
// }

func main() {
	http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/about", about)
	// http.HandleFunc("/contact", contact)
	// http.HandleFunc("/blog", blog)

	fmt.Println("Welcome")
	http.ListenAndServe(":8090", nil)

}
