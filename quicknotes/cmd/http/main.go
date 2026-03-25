package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"quicknotes/internal/handlers"
	"quicknotes/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	config := loadConfig()
	slog.SetDefault(newLogger(os.Stderr, config.GetLevelLog()))

	dbpool, err := pgxpool.New(context.Background(), config.DBConnURL)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer dbpool.Close()

	slog.Info("The connection to the database was successful")

	slog.Info(fmt.Sprintf("Server running on port %s\n", config.ServerPort))
	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("views/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	noteRepo := repositories.NewNoteRepository(dbpool)
	notes, err := noteRepo.List()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	fmt.Println(notes)

	noteHandler := handlers.NewNoteHandler()

	mux.HandleFunc("/", noteHandler.NoteList)
	mux.Handle("/note/view", handlers.HandlerWithError(noteHandler.NoteView))
	mux.HandleFunc("/note/new", noteHandler.NoteNew)
	mux.HandleFunc("/note/create", noteHandler.NoteCreate)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux); err != nil {
		panic(err)
	}
}
