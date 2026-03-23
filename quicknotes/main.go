package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}
type WorldHandler struct{}
type HelloHandler struct{}

func (WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("World"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello"))
}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Rota default")
}

func main() {

	fmt.Println("Server running on port 5000")

	m := MyHandler{}
	mux := http.NewServeMux()
	mux.Handle("/", m)

	http.ListenAndServe(":5000", mux)
}
