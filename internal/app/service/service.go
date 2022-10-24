package service

import (
	"context"

	"github.com/cavelms/internal/app/repository"
	"github.com/gin-gonic/gin"
)

type apiCtx string

var Ctx *gin.Context

type Service interface {
	authService
	userService
	mailService
	fileService
	refereeService
	educationService
	ApiMiddleware(c *gin.Context)
	AuthMidleware(c *gin.Context)
}

type service struct {
	authService
	userService
	mailService
	fileService
	refereeService
	educationService
}

func NewAPIService(repo *repository.Repository) Service {
	return &service{
		authService:      newAuthService(repo),
		userService:      newUserService(repo),
		mailService:      newMailService(repo),
		fileService:      newFileService(repo),
		refereeService:   newRefereeService(repo),
		educationService: newEducationService(repo),
	}
}

func (s *service) ApiMiddleware(c *gin.Context) {
	Ctx = c
	ctxKey := apiCtx("apiCtx")
	ctx := context.WithValue(c.Request.Context(), ctxKey, c)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

func (s *service) AuthMidleware(c *gin.Context) {
	c.Next()
}
