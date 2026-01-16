package utils

import (
	"encoding/json"
	"net/http"
)

// decoding: json ---PARSE---> struct
func ParseBody(r *http.Request, x interface{}) error {
	err := json.NewDecoder(r.Body).Decode(x)
	return err
}

func WriteError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
