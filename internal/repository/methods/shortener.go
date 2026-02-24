package methods

import (
	"URL-Shortener/internal/repository/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type ShortenerMethods struct {
	conn *pgx.Conn
}

func NewShortenerDB(conn *pgx.Conn) *ShortenerMethods {
	return &ShortenerMethods{
		conn: conn,
	}
}

const shortTable = "short_urls"

func (s *ShortenerMethods) Create(url, key string) error {

	return nil
}

func (s *ShortenerMethods) Get(key string) (models.Return, error) {
	var result models.Return

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE key=%s",
		shortTable,
		key,
	)

	rows, err := s.conn.Query(context.Background(), query)
	if err != nil {
		return models.Return{}, nil
	}

	for rows.Next() {
		if err := rows.Scan(
			&result.Key,
			&result.ShortenedLink,
			&result.FullUrl,
			&result.ShortenedLink,
		); err != nil {
			return models.Return{}, err
		}
	}

	return result, nil
}

func (s *ShortenerMethods) Check(key string) bool {
	return false
}
