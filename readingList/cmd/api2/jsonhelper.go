package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (api *Api) writeJSON(w http.ResponseWriter, status int, data any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
func (app *Api) readJSONObject(w http.ResponseWriter, r *http.Request, dto any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var maxBytes int64 = 2_097_152 // 2MB
	http.MaxBytesReader(w, r.Body, maxBytes)

	if err := decoder.Decode(dto); err != nil {
		return err
	}

	err := decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body should  only contain a single object")
	}
	return nil
}
