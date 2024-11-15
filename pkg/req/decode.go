package req

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	defer body.Close()
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	return payload, err
}
