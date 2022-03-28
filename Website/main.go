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
