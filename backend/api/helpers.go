package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *server) respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"data": data,
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, apiErr apiError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiErr.status)
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"error": map[string]interface{}{
			"message": apiErr.message,
			"status":  apiErr.status,
		},
	})
	if err != nil {
		log.Println(err.Error())
	}
	if !apiErr.expected {
		log.Println(apiErr.Error())
	}
}
