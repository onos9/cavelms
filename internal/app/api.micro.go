package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/cavelms/internal/app/repository"
	"github.com/cavelms/internal/app/service"
)

type api struct {
	service service.Service
}

func newAPIService(repo *repository.Repository) *api {
	return &api{
		service: service.NewAPIService(repo),
	}
}

func (api *api) Run(s *http.Server) {
	go func() {
		// service connections
		fmt.Printf("Graphql Server Listening on %s", s.Addr)
		fmt.Println("")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func (api *api) Shutdown(s *http.Server, ctx context.Context) error {
	fmt.Println("\n\nShutting down Grapql service...")
	if err := s.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
