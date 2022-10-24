package app

import (
	"net/http"

	"github.com/cavelms/internal/app/controller"
	"github.com/cavelms/internal/app/repository"
	"github.com/gin-gonic/gin"
)

// const defaultPort = "5050"

type App struct {
	Api    *api
	Server *http.Server
}

func New(r *gin.Engine) *App {
	repo := repository.NewRepository()
	api := newAPIService(repo)
	svr := controller.NewController(r, api.service)
	// fmt.Println(svr)
	return &App{api, svr}
}

func (app *App) Run() {
	app.Api.Run(app.Server)
}
