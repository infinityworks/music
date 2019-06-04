package artist

import (
	"github.com/google/uuid"
	"github.com/infinityworks/music"
)

type Service interface {
	GetAlbums(artistID uuid.UUID) (music.Artist, []music.Album, error)
}

type service struct {
	albumRepository  music.AlbumRepository
	artistRepository music.ArtistRepository
}

func NewService(albumRepository music.AlbumRepository, artistRepository music.ArtistRepository) Service {
	return &service{
		albumRepository:  albumRepository,
		artistRepository: artistRepository,
	}
}

func (r *service) GetAlbums(artistID uuid.UUID) (music.Artist, []music.Album, error) {
	artist, artistErr := r.artistRepository.GetByID(artistID)
	if artistErr != nil {
		return music.Artist{}, []music.Album{}, artistErr
	}

	albums, albumErr := r.albumRepository.GetByArtist(artistID)
	if albumErr != nil {
		return music.Artist{}, []music.Album{}, albumErr
	}

	return artist, albums, nil
}
