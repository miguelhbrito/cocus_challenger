package mhttp

import (
	"encoding/json"
	"net/http"
)

type HttpHandler interface {
	Handler() http.HandlerFunc
}

func WriteJsonResponse(w http.ResponseWriter, data interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		return json.NewEncoder(w).Encode(data)
	}
	return nil
}
