package repositories

import (
	"context"
	"quicknotes/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepository interface {
	List() ([]models.Note, error)
}

type noteRepository struct {
	db *pgxpool.Pool
}

func NewNoteRepository(dbpool *pgxpool.Pool) NoteRepository {
	return &noteRepository{db: dbpool}
}

func (nr *noteRepository) List() ([]models.Note, error) {

	query := `SELECT * FROM quicknotes`
	var list []models.Note

	rows, err := nr.db.Query(context.Background(), query)
	if err != nil {
		return list, nil
	}
	defer rows.Close()

	for rows.Next() {
		var note models.Note
		err = rows.Scan(&note.Id, &note.Title, &note.Content, &note.Author, &note.CreatedAt, &note.UpdateddAt)
		if err != nil {
			return list, nil
		}
		list = append(list, note)
	}

	return list, nil
}
