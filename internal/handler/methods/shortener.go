package methods

import (
	"URL-Shortener/internal/repository/models"
	service "URL-Shortener/internal/services"
	"URL-Shortener/pgk/messages"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortenerMethod struct {
	s service.Shortener
}

func NewShortenerMethod(
	s service.Shortener,
) *ShortenerMethod {
	return &ShortenerMethod{
		s: s,
	}
}

func (sm *ShortenerMethod) POST_reduction(c *gin.Context) {
	var message models.Input
	if err := c.BindJSON(&message); err != nil {
		messages.New(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := sm.s.Get(message.Url)
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error())
		return
	}

	if result == "" {
		r, err := sm.s.Reduction(message.Url)
		if err != nil {
			messages.New(c, http.StatusBadRequest, err.Error())
			return
		}
		result = r
	}

	c.JSON(
		http.StatusOK,
		messages.Message{Message: result},
	)
}

func (sm *ShortenerMethod) GET_redirect(c *gin.Context) {
	key := c.Param("key")

	url, err := sm.s.Get(key)
	if err != nil {
		messages.New(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Redirect(http.StatusFound, url)
}
