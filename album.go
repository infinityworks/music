package music

import "github.com/google/uuid"

//go:generate mockgen -source=album.go -package=mocks -destination=mocks/album_mocks.go

type Album struct {
	Title string `json:"title"`
}

type AlbumRepository interface {
	GetByArtist(artistID uuid.UUID) ([]Album, error)
}
