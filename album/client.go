package album

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/infinityworks/music"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type httpClient struct {
	cli *http.Client
}

func NewRepository() music.AlbumRepository {
	return &httpClient{
		&http.Client{Timeout: time.Second * 5},
	}
}

type albumResponse struct {
	Albums []album `json:"albums"`
}

type album struct {
	Name string `json:"name"`
}

func (r *httpClient) GetByArtist(artistID uuid.UUID) ([]music.Album, error) {
	query := url.Values{}
	query.Add("artist", artistID.String())

	req, _ := http.NewRequest(http.MethodGet, "http://albums.com/albums", nil)
	req.URL.RawQuery = query.Encode()

	res, err := r.cli.Do(req)
	if err != nil {
		return nil, music.ServerError
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, music.ServerError
	}

	var albumResponse albumResponse
	err = json.Unmarshal(data, &albumResponse)
	if err != nil {
		return nil, music.ServerError
	}

	return mapAlbum(albumResponse), nil
}

func mapAlbum(albumResponse albumResponse) []music.Album {
	var albums []music.Album
	for _, a := range albumResponse.Albums {
		albums = append(albums, music.Album{Title: a.Name})
	}
	return albums
}
