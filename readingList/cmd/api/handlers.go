package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (svc *service) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	status := map[string]string{
		"status":      "available",
		"environment": svc.settings.env,
	}

	bytes, _ := json.Marshal(status)
	fmt.Fprintf(w, string(bytes))
}
