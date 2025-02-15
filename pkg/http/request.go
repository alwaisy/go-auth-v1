package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var ErrEmptyBody = errors.New("EMPTY_BODY")

func ParseRequest(r *http.Request, result any) error {
	reqBody := r.Body
	defer func(reqBody io.ReadCloser) {
		err := reqBody.Close()
		if err != nil {
			return
		}
	}(reqBody)

	if reqBody == nil {
		return ErrEmptyBody
	}

	decoder := json.NewDecoder(reqBody)
	decoder.DisallowUnknownFields()

	return decoder.Decode(result)
}
