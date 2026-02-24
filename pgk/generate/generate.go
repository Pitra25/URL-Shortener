package generate

import "math/rand"

func New(length int) (string, error) {
	alphabet := []rune("0123456789abcdefghijklmnopqrstuvwxyz")
	var result = make([]rune, length+1)

	for i := 0; i < length; i++ {
		r := rand.Intn(36)
		result[i] = alphabet[r]
	}

	return string(result), nil
}
