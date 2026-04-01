package repositories

import (
	"context"
	"math/big"
	"quicknotes/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepository interface {
	List() ([]models.Note, error)
	GetById(id int) (*models.Note, error)
	Create(title, content, author string) (*models.Note, error)
	Update(id int, title, content, author string) (*models.Note, error)
	Delete(id int) error
}

type noteRepository struct {
	db *pgxpool.Pool
}

func NewNoteRepository(dbpool *pgxpool.Pool) NoteRepository {
	return &noteRepository{db: dbpool}
}

func (nr *noteRepository) List() ([]models.Note, error) {

	query := `SELECT * FROM notes`
	var list []models.Note

	rows, err := nr.db.Query(context.Background(), query)
	if err != nil {
		return list, nil
	}
	defer rows.Close()

	for rows.Next() {
		var note models.Note
		err = rows.Scan(&note.Id, &note.Title, &note.Content, &note.Author,
			&note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return list, nil
		}
		list = append(list, note)
	}

	return list, nil
}

func (nr *noteRepository) GetById(id int) (*models.Note, error) {

	var note models.Note
	query := `SELECT * FROM notes WHERE id = $1`
	row := nr.db.QueryRow(context.Background(), query, id)

	if err := row.Scan(&note.Id, &note.Title, &note.Content, &note.Author,
		&note.CreatedAt, &note.UpdatedAt); err != nil {
		return &note, err
	}

	return &note, nil
}

func (nr *noteRepository) Create(title, content, author string) (*models.Note, error) {

	var note models.Note
	note.Title = pgtype.Text{String: title, Valid: true}
	note.Content = pgtype.Text{String: content, Valid: true}
	note.Author = pgtype.Text{String: author, Valid: true}

	query := `INSERT INTO notes (title, content, author) 
		VALUES ($1, $2, $3) RETURNING id, created_at`
	row := nr.db.QueryRow(context.Background(), query, note.Title,
		note.Content, note.Author)

	if err := row.Scan(&note.Id, &note.CreatedAt); err != nil {
		return &note, err
	}

	return &note, nil
}

func (nr *noteRepository) Update(id int, title, content, author string) (*models.Note, error) {

	var note models.Note

	note.Id = pgtype.Numeric{Int: big.NewInt(int64(id)), Valid: true}
	if len(title) > 0 {
		note.Title = pgtype.Text{String: title, Valid: true}
	}
	if len(content) > 0 {
		note.Content = pgtype.Text{String: content, Valid: true}
	}
	if len(author) > 0 {
		note.Author = pgtype.Text{String: author, Valid: true}
	}
	note.UpdatedAt = pgtype.Date{Time: time.Now(), Valid: true}

	query := `UPDATE notes title = COALESCE($1, title), content = $2, author = $3,
		updated_at = $4 WHERE id = $5`
	_, err := nr.db.Exec(context.Background(), query, note.Title, note.Content,
		note.Author, note.UpdatedAt, note.Id)

	if err != nil {
		return &note, err
	}

	return &note, nil
}

func (nr *noteRepository) Delete(id int) error {

	query := `DELETE FROM notes WHERE id = $1`
	_, err := nr.db.Exec(context.Background(), query, id)

	if err != nil {
		return err
	}
	return nil
}
