package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/log", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/registration", registration)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8090", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, `welcome`)
	temp, err := template.ParseFiles("template/0_base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	temp.Execute(w, nil)
}
func login(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, `welcome`)
	temp, err := template.ParseFiles("template/0_base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	// template	temp.Execute(w, nil)
	temp, err = temp.ParseFiles("webpage/login.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	temp.Execute(w, nil)
}
func register(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, `welcome`)
	temp, err := template.ParseFiles("template/0_base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	// template	temp.Execute(w, nil)
	temp, err = temp.ParseFiles("webpage/register.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	temp.Execute(w, nil)
}
func registration(w http.ResponseWriter, r *http.Request) {
	// name := r.FormValue("name")
	// email := r.FormValue("mail")
	// pass := r.FormValue("password")
	//display in command line
	//fmt.Println(name,email,pass)
	//display as response in browser
//fmt.Fprintf(w, `succesfully registered name:%s email:%s pass:%s`, name, email, pass)
	//get value via loop 
	r.ParseForm()
	for k, v := range r.Form {
		
		fmt.Println(k,":",v)
	}
	fmt.Fprintln(w, `succesfully registered`)

}
