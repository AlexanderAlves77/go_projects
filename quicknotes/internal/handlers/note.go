package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"quicknotes/internal/apperror"
)

type noteHandler struct{}

func NewNoteHandler() *noteHandler {
	return &noteHandler{}
}

func (nh *noteHandler) NoteList(w http.ResponseWriter, r *http.Request) {

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

	slog.Info("Executou o handler / ")
	temp.ExecuteTemplate(w, "base", nil)
}

func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) error {

	id := r.URL.Query().Get("id")
	if id == "" {
		return apperror.WithStatus(errors.New("anotação é obrigatória"),
			http.StatusBadRequest)
	}

	if id == "0" {
		//return handlers.ErrNotFound
		return apperror.WithStatus(errors.New("anotação 0 não foi encontrada"), http.StatusNotFound)
	}

	files := []string{
		"views/templates/base.html",
		"views/pages/noteView.html",
	}
	temp, err := template.ParseFiles(files...)

	if err != nil {
		return errors.New("aconteceu um erro ao executar essa página")
	}

	return temp.ExecuteTemplate(w, "base", id)
}

func (nh *noteHandler) NoteNew(w http.ResponseWriter, r *http.Request) {

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

func (nh *noteHandler) NoteCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Método não permitido", 405)
		return
	}
	fmt.Fprint(w, "Criando uma nova nota...")
}
