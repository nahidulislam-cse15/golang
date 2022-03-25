package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

type Todo struct {
	Id int
	Name   string
	Status string
	Day   int
	Month time.Month
}

func main() {

	text := `your task{{.Id}} is {{ .Name}} and the task  {{if eq .Status "yes"}} is done {{else}} is left{{end}} Today is: {{ .Day }} {{.Month }}
`
	current_time := time.Now()
	task1 := Todo{1,"golang text template", "yes",current_time.Day(),current_time.Month()}
	task2 := Todo{2,"complete english module 2","yes",current_time.Day(),current_time.Month()}
	task3 := Todo{3,"complete english module 3","no",current_time.Day(),current_time.Month()}
	//fmt.Println(work.name)
	todos := template.New("todos")

	t, err := todos.Parse(text)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = t.Execute(os.Stdout, task1)
	if err != nil {
		fmt.Println(err.Error)
	}
	err = t.Execute(os.Stdout, task2)
	if err != nil {
		fmt.Println(err.Error)
	}
	err = t.Execute(os.Stdout, task3)
	if err != nil {
		fmt.Println(err.Error)
	}


}

