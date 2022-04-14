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
// //year month day
// year, month, day := time.Now().Date()
// // Create a new file in the uploads directory
// dst, err := os.Create(fmt.Sprintf("./uploads/%d-%d-%d%s", year, int(month), day, filepath.Ext(fileHeader.Filename)))
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	return
// }
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 32 MB is the default used by FormFile()
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get a reference to the fileHeaders.
	// They are accessible only after ParseMultipartForm is called
	files := r.MultipartForm.File["file"]

	for _, fileHeader := range files {
		// Restrict the size of each uploaded file to 1MB.
		// To prevent the aggregate size from exceeding
		// a specified value, use the http.MaxBytesReader() method
		// before calling ParseMultipartForm()
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}

		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
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
