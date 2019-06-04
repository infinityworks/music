package api

import (
	"github.com/infinityworks/music"
	"net/http"

	"github.com/gin-gonic/gin"
)

// errorToStatusCodeLookup maps application errors to http status codes
// It helps to decouple application errors from the HTTP layer
var errorToStatusCodeLookup = map[string]int{
	music.InvalidID.Code:    http.StatusBadRequest,
	music.UserNotFound.Code: http.StatusBadRequest,
	music.ServerError.Code:  http.StatusInternalServerError,
}

// errorHandler is a middleware that sets any present application errors on the response
func errorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	for _, err := range c.Errors {
		if parsedError, ok := err.Err.(music.Error); ok {
			statusCode := errorToStatusCodeLookup[parsedError.Code]
			c.JSON(statusCode, parsedError)
			return
		}
	}

	c.JSON(http.StatusInternalServerError, music.ServerError)
}
