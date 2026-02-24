package handler

import (
	"URL-Shortener/internal/handler/methods"
	service "URL-Shortener/internal/services"

	"github.com/gin-gonic/gin"
)

type (
	Auth interface {
		Register(c *gin.Context)
		Login(c *gin.Context)
	}

	Shortener interface {
		POST_reduction(c *gin.Context)
		GET_redirect(c *gin.Context)
	}

	HandlerStr struct {
		Auth
		Shortener
	}
)

func NewHandlers(service *service.ServiceStr) *HandlerStr {
	return &HandlerStr{
		Auth:      methods.NewAuthMethod(service.Auth),
		Shortener: methods.NewShortenerMethod(service.Shortener),
	}
}
