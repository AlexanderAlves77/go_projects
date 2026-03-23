package main

import (
	"fmt"
	"html/template"
	"os"
)

type TemplateData struct {
	Nome string
	Age  int
}

func main() {
	template, err := template.ParseFiles("hello.html")
	fmt.Println(template.Name())
	if err != nil {
		panic(err)
	}

	data := TemplateData{Nome: "Alexander"}
	err = template.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
