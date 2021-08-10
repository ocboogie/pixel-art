package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"

	log "github.com/sirupsen/logrus"
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
		log.Errorln(err.Error())
	}
	if !apiErr.expected {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Errorf("Could not create request dump: %v", err)
		}

		dumpString := fmt.Sprintf("%q", dump)
		log.WithField("requestDump", dumpString).Errorf("Unexpected API Error: %v", apiErr)
	}
}
