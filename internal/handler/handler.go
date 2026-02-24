package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	m *HandlerStr
}

func New(m *HandlerStr) *Handler {
	return &Handler{m: m}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", h.m.Register)
			auth.POST("/login", h.m.Login)
		}

		api.POST("/example", h.m.POST_reduction)
		api.GET("/:key", h.m.GET_redirect)
	}

	return router
}
