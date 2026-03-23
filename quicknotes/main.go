package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}
type WorldHandler struct{}
type HelloHandler struct{}

func noteList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header()["Date"] = nil

	fmt.Fprint(w, `<h1>Lista de anotações e lembretes</h1>`)
}

func noteView(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Nota não encontrada", http.StatusNotFound)
		return
	}
	note := `
		<div>
			<h3>Está é a nota 1</h3>
			<p>Este é o conteúdo da anotação</p>
		</div>
	`
	fmt.Fprint(w, note)
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
