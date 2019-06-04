package sqlite_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/infinityworks/music"
	"github.com/infinityworks/music/sqlite"
	"github.com/infinityworks/music/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetByID_Success(t *testing.T) {
	id := uuid.New()
	db := test.DBSetup(func(tx *sqlx.Tx) {
		tx.MustExec(fmt.Sprintf(
			`INSERT INTO artists (external_id, name) VALUES ('%s', 'Sika Redem')`, id))
	})
	repo := sqlite.NewArtistRepository(db)

	artist, err := repo.GetByID(id)

	assert.NoError(t, err)
	assert.Equal(t, music.Artist{Name: "Sika Redem"}, artist)
}

func TestGetByID_NotFound(t *testing.T) {
	id := uuid.New()
	db := test.DBSetup(func(tx *sqlx.Tx) {})
	repo := sqlite.NewArtistRepository(db)

	_, err := repo.GetByID(id)

	assert.EqualError(t, err, music.UserNotFound.Code)
}

func TestGetByID_ServerErrorIfCannotMapError(t *testing.T) {
	id := uuid.New()
	db := test.DBSetup(func(tx *sqlx.Tx) {})
	err := db.Close()
	if err != nil {
		panic(err)
	}
	repo := sqlite.NewArtistRepository(db)

	_, err = repo.GetByID(id)

	assert.EqualError(t, err, music.ServerError.Code)
}