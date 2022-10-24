package controller

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cavelms/config"
	"github.com/cavelms/graph/generated"
	"github.com/cavelms/internal/app/service"
	"github.com/gin-gonic/gin"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.Service
}

func NewController(r *gin.Engine, s service.Service) *http.Server {
	re := &Resolver{Service: s}
	gqlcfg := generated.Config{Resolvers: re}

	// r.Use()
	r.GET("/", re.playgroundHanler())
	r.POST("/query", s.ApiMiddleware, re.queryHanler(gqlcfg))
	r.POST("/upload", s.AuthMidleware, re.uploadHandler)

	config := config.GetConfig()
	port := fmt.Sprintf(":%s", config.Port)

	return &http.Server{
		Addr:    port,
		Handler: r,
	}
}

func (r *Resolver) queryHanler(cfg generated.Config) gin.HandlerFunc {
	cfg.Directives.RequireAuth = r.Service.RequireAuth
	schema := generated.NewExecutableSchema(cfg)
	h := handler.NewDefaultServer(schema)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (r *Resolver) playgroundHanler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (r *Resolver) uploadHandler(c *gin.Context) {

}
