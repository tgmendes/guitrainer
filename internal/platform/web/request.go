package web

import (
	"encoding/json"
	"net/http"
)

func Decode(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(data); err != nil {
		return err
	}

	return nil
}
