package main

import (
	"fmt"
	"os"
	"text/template"
	//"time"
)

func main() {

template.
	temp, err:= template.ParseFiles("one.gohtml","two.gohtml") //return first files
	if err != nil {
		fmt.Println(err.Error)
	}
	err = temp.Execute(os.Stdout, "nahid") //2nd param may b nil or any type received by {{.}} (pipelining)


}

