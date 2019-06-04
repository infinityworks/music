package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/infinityworks/music"
	"net/http"
)

type response struct {
	Artist music.Artist  `json:"artist"`
	Albums []music.Album `json:"albums"`
}

func (r *Server) getArtistAlbums(ctx *gin.Context) {
	id := ctx.Param("id")
	externalID, err := uuid.Parse(id)
	if err != nil {
		ctx.Error(music.InvalidID)
		return
	}

	artist, albums, err := r.service.GetAlbums(externalID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response{
		Artist: artist,
		Albums: albums,
	})
}
