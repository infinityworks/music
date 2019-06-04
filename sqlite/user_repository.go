package sqlite

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/infinityworks/music"
	"github.com/jmoiron/sqlx"
)

type artistRepository struct {
	db *sqlx.DB
}

func NewArtistRepository(db *sqlx.DB) music.ArtistRepository {
	return &artistRepository{db}
}

func (r *artistRepository) GetByID(id uuid.UUID) (music.Artist, error) {
	var name string
	err := r.db.Get(&name, `SELECT name FROM artists WHERE external_id = $1`, id)
	if err != nil {
		return music.Artist{}, mapError(err)
	}
	return music.Artist{Name: name}, nil
}

func mapError(err error) error {
	if err == sql.ErrNoRows {
		return music.UserNotFound
	}
	return music.ServerError
}
