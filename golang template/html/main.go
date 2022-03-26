package main

import (
	"fmt"
	"os"
	"text/template"
	//"time"
)

func main() {


	//temp, err:= template.ParseFiles("one.gohtml","two.gohtml") //return first files
	temp,err := template.ParseGlob("*.gohtml")//increasing order filename selected
	if err != nil {
		fmt.Println(err.Error)
	}
	//data pass
	names:=[]string{"nahid","tasin","tasrif","raiyan"}
	err = temp.Execute(os.Stdout, names) //2nd param may b nil or any type received by {{.}} (pipelining)
     
//	err=temp.ExecuteTemplate(os.Stdout,"two.gohtml",nil)

}

