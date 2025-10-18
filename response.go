package chikit

import (
	"encoding/json"
	"net/http"

	"github.com/opentdp/go-helper/logman"
)

// SendJSON sends a JSON response with the given status code and data
func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logman.Error("failed to encode JSON response", "error", err)
	}
}

// SendSuccess sends a success response with default 200 status
func SendSuccess(w http.ResponseWriter, data interface{}) {
	SendJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

// SendError sends an error response
func SendError(w http.ResponseWriter, statusCode int, message string) {
	SendJSON(w, statusCode, map[string]interface{}{
		"success": false,
		"message": message,
	})
}
