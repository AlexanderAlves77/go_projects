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

	temp, err := template.ParseFiles("views/templates/home.html")
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}

	temp.Execute(w, nil)
}

func noteView(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Nota não encontrada", http.StatusNotFound)
		return
	}

	temp, err := template.ParseFiles("views/templates/noteView.html")
	if err != nil {
		http.Error(w, "Aconteceu um erro ao executar essa página", http.StatusInternalServerError)
		return
	}

	temp.Execute(w, id)
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
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":5000", mux)
}
