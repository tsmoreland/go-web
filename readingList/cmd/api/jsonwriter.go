package main

import (
	"encoding/json"
	"net/http"
)

func (svc *service) writeJSON(w http.ResponseWriter, status int, data any) error {
	if js, err := json.Marshal(data); err != nil {
		return err
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(js)
		return nil
	}
}
