package methods

import (
	database "URL-Shortener/internal/repository"
	"URL-Shortener/pgk/generate"
	"fmt"
)

type ShortenerService struct {
	repo database.Shortener
}

func NewShortenerService(repo database.Shortener) *ShortenerService {
	return &ShortenerService{
		repo: repo,
	}
}

/*

 url = https://host/...

*/

func (s *ShortenerService) Reduction(link string) (string, error) {

	var (
		maxAttempts = 15
		newKey      = ""
		err         error
	)

	for i := 0; i < maxAttempts; i++ {
		newKey, err = generate.New(7)
		if err != nil {
			return "", err
		}

		if !s.repo.Check(newKey) {
			break
		}

		if i == maxAttempts-1 {
			return "", fmt.Errorf("Failed to generate unique key")
		}
	}

	err = s.repo.Create(link, newKey)
	if err != nil {
		return "", err
	}

	return newKey, nil
}

func (s *ShortenerService) Get(key string) (string, error) {
	return "", nil
}
