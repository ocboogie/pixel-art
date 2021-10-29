package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func paramExists(r *http.Request, param string) bool {
	if err := r.ParseForm(); err != nil {
		return false
	}

	_, exists := r.Form[param]
	return exists
}

func paramNumber(r *http.Request, param string) (int, error) {
	paramValue := r.URL.Query().Get(param)
	i, err := strconv.Atoi(paramValue)

	if err != nil {
		return 0, err
	}

	return i, nil
}

func paramTime(r *http.Request, param string) (*time.Time, error) {
	paramValue := r.URL.Query().Get(param)
	paramTime, err := time.Parse(time.RFC3339, paramValue)

	if err != nil {
		return nil, err
	}

	return &paramTime, nil
}

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
