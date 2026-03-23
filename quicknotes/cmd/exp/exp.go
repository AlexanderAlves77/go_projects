package main

import (
	"html/template"
	"os"
)

type TemplateData struct {
	Nome string
	Age  int
}

func main() {
	template, err := template.ParseFiles("layout1.html", "layout2.html", "home.html", "footer.html", "header.html")

	//fmt.Println(template.Name())
	//fmt.Println(template.DefinedTemplates())

	if err != nil {
		panic(err)
	}

	//data := TemplateData{Nome: "Alexander"}
	err = template.ExecuteTemplate(os.Stdout, "layout1.html", "2026")
	if err != nil {
		panic(err)
	}
}
