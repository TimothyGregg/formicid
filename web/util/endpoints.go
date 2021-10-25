package util

import (
	"net/http"
)

type Endpoint struct {
	http.Handler
	Default http.HandlerFunc
	Post    http.HandlerFunc
	Get     http.HandlerFunc
	Put     http.HandlerFunc
	Patch   http.HandlerFunc
	Delete  http.HandlerFunc
}

func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if e.Post != nil {
			e.Post(w, r)
			return
		}
	case http.MethodGet:
		if e.Get != nil {
			e.Get(w, r)
			return
		}
	case http.MethodPut:
		if e.Put != nil {
			e.Put(w, r)
			return
		}
	case http.MethodPatch:
		if e.Patch != nil {
			e.Patch(w, r)
			return
		}
	case http.MethodDelete:
		if e.Delete != nil {
			e.Delete(w, r)
			return
		}
	default:
		if e.Default != nil {
			e.Default(w, r)
			return
		}
	}
	ErrorResponse(w, r.Method + " not accepted at this endpoint", http.StatusMethodNotAllowed)
}
