package api

import (
	"context"
	"github.com/infinityworks/music/artist"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	service artist.Service
	server  *http.Server
}

// NewServer creates a new server with all application routes defined
// The caller must call `Start` to bind to the network and start serving requests
func NewServer(service artist.Service) *Server {
	r := gin.Default()

	srv := &Server{Router: r, service: service}
	r.GET("/health", srv.healthCheck)

	v1 := r.Group("/v1")
	v1.Use(errorHandler)

	v1.GET("/artists/:id/albums", srv.getArtistAlbums)

	return srv
}

func (r *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func (r *Server) Start(port string) {
	r.server = &http.Server{
		Addr:    ":" + port,
		Handler: r.Router,
	}
	defer r.Close()

	go func() {
		if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func (r *Server) Close() {
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := r.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown failed: %s.", err)
	}

	log.Println("Server shutdown complete.")
}
