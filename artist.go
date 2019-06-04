package music

import "github.com/google/uuid"

//go:generate mockgen -source=artist.go -package=mocks -destination=mocks/artist_mocks.go

type Artist struct {
	Name string `json:"name"`
}

type ArtistRepository interface {
	GetByID(id uuid.UUID) (Artist, error)
}
