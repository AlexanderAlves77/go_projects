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

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

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

	temp.ExecuteTemplate(w, "base", id)
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

	config := loadConfig()

	fmt.Printf("Server running on port %s", config.ServerPort)
	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("views/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/new", noteNew)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux)
}
