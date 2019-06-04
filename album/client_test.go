package album_test

import (
	"github.com/google/uuid"
	"github.com/infinityworks/music"
	"github.com/infinityworks/music/album"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetByArtist_WithAlbums(t *testing.T) {
	albumClient := album.NewRepository()
	id := uuid.MustParse("43e0b4a6-a588-46b4-8969-1efa3a4434ee")
	defer apitest.NewMock().
		Get("/albums").
		Query("artist", id.String()).
		RespondWith().
		Status(http.StatusOK).
		Body(`{"albums": [{"name": "Entheogen"}]}`).
		EndStandalone()()

	albums, err := albumClient.GetByArtist(id)

	assert.NoError(t, err)
	assert.Equal(t, []music.Album{{Title: "Entheogen"}}, albums)
}
