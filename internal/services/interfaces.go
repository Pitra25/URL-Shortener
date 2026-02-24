package service

import (
	database "URL-Shortener/internal/repository"
	"URL-Shortener/internal/services/methods"
)

type (
	Auth interface {
		// Add necessary fields here
	}

	Shortener interface {
		Reduction(link string) (string, error)
		Get(key string) (string, error)
	}

	ServiceStr struct {
		Auth
		Shortener
	}
)

func New(repo *database.Repository) *ServiceStr {
	return &ServiceStr{
		Shortener: methods.NewShortenerService(repo.Shortener),
	}
}
