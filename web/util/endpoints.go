package util

import (
	"net/http"
)

type Endpoint struct {
	http.Handler
	Default http.Handler
	Post    http.Handler
	Get     http.Handler
	Put     http.Handler
	Patch   http.Handler
	Delete  http.Handler
}

func NewEndpoint(function http.Handler) *Endpoint {
	return &Endpoint{Default: function}
}

/*
func (e *Endpoint) AddHandler(method string, function http.HandlerFunc) error {
	switch method {
	case http.MethodPost:
		e.Post = function
	case http.MethodGet:
		e.Get = function
	case http.MethodPut:
		e.Put = function
	case http.MethodPatch:
		e.Patch = function
	case http.MethodDelete:
		e.Delete = function
	default:
		return errors.New("this wasn't a valid method, dingus")
	}
	return nil
}
*/

func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if e.Post != nil {
			e.Post.ServeHTTP(w, r)
			return
		}
	case http.MethodGet:
		if e.Get != nil {
			e.Get.ServeHTTP(w, r)
			return
		}
	case http.MethodPut:
		if e.Put != nil {
			e.Put.ServeHTTP(w, r)
			return
		}
	case http.MethodPatch:
		if e.Patch != nil {
			e.Patch.ServeHTTP(w, r)
			return
		}
	case http.MethodDelete:
		if e.Delete != nil {
			e.Delete.ServeHTTP(w, r)
			return
		}
	default:
		if e.Default != nil {
			e.Default.ServeHTTP(w, r)
			return
		}
	}
	ErrorResponse(w, r.Method + " not accepted at this endpoint", http.StatusMethodNotAllowed)
}
