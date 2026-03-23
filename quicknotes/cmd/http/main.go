package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type MyHandler struct{}
type WorldHandler struct{}
type HelloHandler struct{}

func noteList(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"views/templates/base.html",
		"views/pages/home.html",
	}
	temp, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}

	temp.ExecuteTemplate(w, "base", nil)
}

func noteView(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Nota não encontrada", http.StatusNotFound)
		return
	}

	files := []string{
		"views/templates/base.html",
		"views/pages/noteView.html",
	}
	temp, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}

	temp.ExecuteTemplate(w, "base", nil)
}

func noteNew(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"views/templates/base.html",
		"views/pages/noteNew.html",
	}
	temp, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}

	temp.ExecuteTemplate(w, "base", nil)
}

func noteCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Método não permitido", 405)
		return
	}
	fmt.Fprint(w, "Criando uma nova nota...")
}

func main() {

	fmt.Println("Server running on port 5000")
	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/new", noteNew)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":5000", mux)
}
