package methods

import (
	service "URL-Shortener/internal/services"
	"URL-Shortener/pgk/messages"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMethods struct {
	s service.Auth
}

func NewAuthMethod(
	s service.Auth,
) *AuthMethods {
	return &AuthMethods{
		s: s,
	}
}

func (am *AuthMethods) Register(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		messages.Message{
			Message: "Good registration",
		},
	)
}

func (am *AuthMethods) Login(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		messages.Message{
			Message: "Good login",
		},
	)
}
