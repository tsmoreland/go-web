package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (api *Api) writeJSON(w http.ResponseWriter, status int, data any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
func (api *Api) readJSONObject(w http.ResponseWriter, r *http.Request, dto any) error {
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

func (api *Api) writeProblemDetails(w http.ResponseWriter, r *http.Request, title string,
	statusCode int, detail string) {
	problem := &ProblemDetails{
		Type_:    fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/%d", statusCode),
		Title:    title,
		Status:   int32(statusCode),
		Detail:   detail,
		Instance: r.URL.String(),
	}

	if err := api.writeJSON(w, statusCode, problem); err != nil {
		api.logger.Print(err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
