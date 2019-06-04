package api_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/infinityworks/music/album"
	"github.com/infinityworks/music/api"
	"github.com/infinityworks/music/artist"
	"github.com/infinityworks/music/sqlite"
	"github.com/infinityworks/music/test"
	"github.com/jmoiron/sqlx"
	"github.com/steinfletcher/apitest"
	"net/http"
	"testing"
)

func TestGetArtistAlbums_Success(t *testing.T) {
	id := uuid.New()

	test.DBSetup(func(tx *sqlx.Tx) {
		tx.MustExec(fmt.Sprintf(
			`INSERT INTO artists (external_id, name) VALUES ('%s', 'Sika Redem')`, id))
	})

	apiTest().
		Mocks(getAlbumsMock(id)).
		Get(fmt.Sprintf("/v1/artists/%s/albums", id)).
		Expect(t).
		Body(`{"artist":{"name":"Sika Redem"},"albums":[{"title":"Entheogen"}]}`).
		Status(http.StatusOK).
		End()
}

func TestGetArtistAlbums_UserNotFound(t *testing.T) {
	apiTest().
		Get(fmt.Sprintf("/v1/artists/%s/albums", uuid.New())).
		Expect(t).
		Body(`{"code":"USER_NOT_FOUND", "detail":"Artist not found"}`).
		Status(http.StatusBadRequest).
		End()
}

func TestGetArtistAlbums_AlbumApiError(t *testing.T) {
	id := uuid.New()

	test.DBSetup(func(tx *sqlx.Tx) {
		tx.MustExec(fmt.Sprintf(
			`INSERT INTO artists (external_id, name) VALUES ('%s', 'Sika Redem')`, id))
	})

	apiTest().
		Mocks(apitest.NewMock().
			Get("http://albums.com/albums").
			RespondWith().
			Status(http.StatusInternalServerError).
			End()).
		Get(fmt.Sprintf("/v1/artists/%s/albums", id)).
		Expect(t).
		Body(`{"code":"SERVER_ERROR", "detail":"Sorry, something went wrong"}`).
		Status(http.StatusInternalServerError).
		End()
}

func getAlbumsMock(id uuid.UUID) *apitest.Mock {
	return apitest.NewMock().
		Get("http://albums.com/albums").
		Query("artist", id.String()).
		RespondWith().
		Body(`{"albums":[{"name": "Entheogen"}]}`).
		Status(http.StatusOK).
		End()
}

func apiTest() *apitest.APITest {
	db := test.DBConnect()
	service := artist.NewService(album.NewRepository(), sqlite.NewArtistRepository(db))
	return apitest.New().
		Recorder(test.Recorder).
		Report(apitest.SequenceDiagram()).
		Handler(api.NewServer(service).Router)
}
