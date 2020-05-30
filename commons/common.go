package commons

import (
	"encoding/json"
	"net/http"
)

// RespondJSON makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, code int, payload interface{}) {
	r, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(r))
}

// RespondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}
