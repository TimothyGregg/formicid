package util

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// 204 No Content
func Response_NoContent(w http.ResponseWriter, headerContent map[string]string) {
	w.Header().Set("Date", time.Now().UTC().Format(http.TimeFormat))
	for field, value := range headerContent {
		w.Header().Set(field, value)
	}
	w.WriteHeader(http.StatusNoContent)
}

// 400 Bad Request
func Response_BadRequest(w http.ResponseWriter, additionalMessages ...string) {
	// Header
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Date", time.Now().UTC().Format(http.TimeFormat))
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusBadRequest)

	// Body
	resp := make(map[string]string)
	resp["error"] = "Bad Request"
	if len(additionalMessages) > 0 {
		resp["additional"] = strings.Join(additionalMessages, ", ")
	}
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

// 404 Not Found
func Response_NotFound(w http.ResponseWriter, additionalMessages ...string) {
	// Header
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Date", time.Now().UTC().Format(http.TimeFormat))
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusNotFound)

	// Body
	resp := make(map[string]string)
	resp["error"] = "Not Found"
	if len(additionalMessages) > 0 {
		resp["additional"] = strings.Join(additionalMessages, ", ")
	}
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

// 405 Method Not Allowed
func Response_MethodNotAllowed(w http.ResponseWriter, allowedMethods []string) {
	// Header]\]\]\\
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Date", time.Now().UTC().Format(http.TimeFormat))
	w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
	w.WriteHeader(http.StatusMethodNotAllowed)

	// Body
	resp := make(map[string]string)
	resp["error"] = "Method Not Allowed"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

// 415 Unsupported Media Type
func Response_UnsupportedMediaType(w http.ResponseWriter, additionalMessages ...string) {
	// Header
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Date", time.Now().UTC().Format(http.TimeFormat))
	w.WriteHeader(http.StatusUnsupportedMediaType)

	// Body
	resp := make(map[string]string)
	resp["error"] = "Unsupported Media Type"
	if len(additionalMessages) > 0 {
		resp["additional"] = strings.Join(additionalMessages, ", ")
	}
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

// 503 Service Unavailable
func Response_ServerUnavailable(w http.ResponseWriter, additionalMessages ...string) {
	// Header
	w.Header().Set("Content-Type", "application/")
	w.Header().Set("Date", time.Now().UTC().Format(http.TimeFormat))
	w.WriteHeader(http.StatusServiceUnavailable)

	// Body
	resp := make(map[string]string)
	resp["error"] = "Service Unavailable"
	if len(additionalMessages) > 0 {
		resp["additional"] = strings.Join(additionalMessages, ", ")
	}
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
