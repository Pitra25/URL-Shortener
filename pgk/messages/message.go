package messages

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type (
	Message struct {
		Message string `json:"message"`
	}

	StatusMessage struct {
		Status string `json:"status"`
	}
)

func New(c *gin.Context, statusCode int, message string) {
	logrus.Error(
		"Status: ", statusCode,
		" Message: ", message,
	)

	c.AbortWithStatusJSON(statusCode, message)
}
