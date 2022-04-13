package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

func home(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(w.Header().Get("Content-Type"))
	//	fmt.Fprint(w, "<h1>Welcome to Master Academy</h1>")
	//fmt.Println(w.Response.StatusCode)
	//http.StatusNotFound
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {

		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)

		return

	}
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//year month day
	year, month, day := time.Now().Date()
	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./uploads/%d-%d-%d%s", year, int(month), day, filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Upload successful")

}
func pathHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, r.URL.Path)    //deoceded path
	fmt.Fprintln(w, r.URL.RawPath) //encoded path
	switch r.URL.Path {
	case "/":
		home(w, r)
	case "/about":
		about(w, r)
	case "/contact":
		contact(w, r)
	case "/blog":
		blog(w, r)
	default:
		//way 1
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w,"Page not found")

		//way 2
		http.Error(w, "Page not found", http.StatusNotFound)
		//way 3
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

	}

}

func about(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<h1>Welcome to Master Academy about page</h1>")
}
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to Master Academy contact page</h1>")
}
func blog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to Master Academy blog page</h1>")
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		home(w, r)
	case "/about":
		about(w, r)
	case "/contact":
		contact(w, r)
	case "/blog":
		blog(w, r)
	default:
		//way 1
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w,"Page not found")

		//way 2
		http.Error(w, "Page not found", http.StatusNotFound)
		//way 3
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

	}
}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/upload", uploadHandler)
	//http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/about", about)
	// http.HandleFunc("/contact", contact)
	// http.HandleFunc("/blog", blog)
	//http.Handle("/",http.HandlerFunc(pathHandler))
	//http.HandleFunc("/",http.HandlerFunc(pathHandler).ServeHTTP)//default servemux

	fmt.Println("Welcome")
	http.ListenAndServe(":8030", nil)
	// var router http.HandlerFunc
	// router=pathHandler
	//http.ListenAndServe(":8090", router)
	//http.ListenAndServe(":8030", http.HandlerFunc(pathHandler)) //pathhandler func converted to handler type

}
