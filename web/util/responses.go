package util

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func ok(w http.ResponseWriter) {
	ErrorResponse(w, "Success", http.StatusOK)
}
