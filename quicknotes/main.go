package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}
type WorldHandler struct{}
type HelloHandler struct{}

func noteList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header()["Date"] = nil

	fmt.Fprint(w, `{"id": 1}`)
}

func noteView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Exibindo uma nota...")
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
