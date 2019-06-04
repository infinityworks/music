package artist_test

import (
	"github.com/google/uuid"
	"github.com/infinityworks/music"
	"github.com/infinityworks/music/artist"
	"github.com/infinityworks/music/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAlbums_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	albumRepo := mocks.NewMockAlbumRepository(ctrl)
	artistRepo := mocks.NewMockArtistRepository(ctrl)
	service := artist.NewService(albumRepo, artistRepo)
	id := uuid.MustParse("2995181a-8bd9-4960-91e4-a7ada06e48a3")
	anAlbum := anAlbum()
	anArtist := anArtist()

	albumRepo.EXPECT().GetByArtist(gomock.Eq(id)).Return(anAlbum, nil)
	artistRepo.EXPECT().GetByID(gomock.Eq(id)).Return(anArtist, nil)

	actualArtist, actualAlbums, err := service.GetAlbums(id)

	assert.NoError(t, err)
	assert.Equal(t, anArtist, actualArtist)
	assert.Equal(t, anAlbum, actualAlbums)
}

func TestGetAlbums_AlbumErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	albumRepo := mocks.NewMockAlbumRepository(ctrl)
	artistRepo := mocks.NewMockArtistRepository(ctrl)
	service := artist.NewService(albumRepo, artistRepo)
	id := uuid.MustParse("2995181a-8bd9-4960-91e4-a7ada06e48a3")
	anArtist := anArtist()

	albumRepo.EXPECT().GetByArtist(gomock.Eq(id)).Return([]music.Album{}, errors.New("failed to fetch albums"))
	artistRepo.EXPECT().GetByID(gomock.Eq(id)).Return(anArtist, nil)

	_, _, err := service.GetAlbums(id)

	assert.EqualError(t, err, "failed to fetch albums")
}

func TestGetAlbums_ArtistErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	artistRepo := mocks.NewMockArtistRepository(ctrl)
	service := artist.NewService(nil, artistRepo)
	id := uuid.MustParse("2995181a-8bd9-4960-91e4-a7ada06e48a3")

	artistRepo.EXPECT().GetByID(gomock.Eq(id)).Return(music.Artist{}, errors.New("failed to fetch artist"))

	_, _, err := service.GetAlbums(id)

	assert.EqualError(t, err, "failed to fetch artist")
}

func anArtist() music.Artist {
	return music.Artist{
		Name: "Sika Redem",
	}
}

func anAlbum() []music.Album {
	return []music.Album{
		{
			Title: "Entheogen",
		},
	}
}