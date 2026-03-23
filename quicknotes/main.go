package main

import (
	"fmt"
	"net/http"
)

type WorldHandler struct{}
type HelloHandler struct{}

func (WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("World"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello"))
}

func main() {

	fmt.Println("Server running on port 5000")

	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":5000", nil)
}
