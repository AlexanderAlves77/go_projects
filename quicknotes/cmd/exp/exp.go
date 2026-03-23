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
	template, err := template.ParseFiles("a.html", "b.html")
	fmt.Println(template.Name())
	if err != nil {
		panic(err)
	}

	//data := TemplateData{Nome: "Alexander"}
	err = template.ExecuteTemplate(os.Stdout, "a", nil)
	if err != nil {
		panic(err)
	}
}
