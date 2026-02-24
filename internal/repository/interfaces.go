package database

import (
	"URL-Shortener/internal/repository/methods"
	"URL-Shortener/internal/repository/models"

	"github.com/jackc/pgx/v5"
)

type (
	Auth interface {
		// Add method signatures here
	}

	Shortener interface {
		Create(url, key string) error
		Get(key string) (models.Return, error)
		Check(key string) bool
	}

	Repository struct {
		Auth
		Shortener
	}
)

func New(conn *pgx.Conn) *Repository {
	return &Repository{
		Auth:      methods.NewAuthDB(conn),
		Shortener: methods.NewShortenerDB(conn),
	}
}
