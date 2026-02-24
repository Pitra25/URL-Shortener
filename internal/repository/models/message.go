package models

import "fmt"

type (
	Input struct {
		Url string `json:"url"`
	}

	Response struct {
		ShortenedLink string `json:"shortenedLink"`
	}

	Return struct {
		Key           string `json:"key"`
		ShortenedLink string `json:"short_url"`
		FullUrl       string `json:"full_url"`
		CreatedAt     string `json:"created_at"`
	}
)

func (i *Input) Validate() error {
	if i.Url != "" {
		return fmt.Errorf("input structure has no values")
	}
	return nil
}
