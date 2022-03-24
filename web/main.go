package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(w.Header().Get("Content-Type"))
	fmt.Fprint(w, "<h1>Welcome to Master Academy</h1>")
	//fmt.Println(w.Response.StatusCode)
	//http.StatusNotFound
}

func pathHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, r.URL.Path)//deoceded path
	fmt.Fprintln(w, r.URL.RawPath)//encoded path
	switch r.URL.Path{
	 case "/":
		home(w,r)
	case "/about":
		about(w,r)	
	case "/contact":
		contact(w,r)
	case "/blog":
		blog(w,r)	
	default:
		//way 1
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w,"Page not found")

		//way 2
		http.Error(w,"Page not found",http.StatusNotFound)
		//way 3
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)


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
type Router struct {} 
func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path{
	case "/":
	   home(w,r)
   case "/about":
	   about(w,r)	
   case "/contact":
	   contact(w,r)
   case "/blog":
	   blog(w,r)	
   default:
	   //way 1
	   // w.WriteHeader(http.StatusNotFound)
	   // fmt.Fprint(w,"Page not found")

	   //way 2
	   http.Error(w,"Page not found",http.StatusNotFound)
	   //way 3
	   http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)


   }
}
func main() {

	//http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/about", about)
	// http.HandleFunc("/contact", contact)
	// http.HandleFunc("/blog", blog)
	//http.Handle("/",http.HandlerFunc(pathHandler))
	//http.HandleFunc("/",http.HandlerFunc(pathHandler).ServeHTTP)//default servemux

	fmt.Println("Welcome")
	//http.ListenAndServe(":8090", nil)
	// var router http.HandlerFunc
	// router=pathHandler
	//http.ListenAndServe(":8090", router)
	http.ListenAndServe(":8090", http.HandlerFunc(pathHandler))//pathhandler func converted to handler type


}
