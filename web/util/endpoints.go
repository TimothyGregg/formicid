package util

import (
	"errors"
	"net/http"
	"strings"
)

var	HTMLMethods = [6]string{
		http.MethodOptions,
		http.MethodPost,
		http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodDelete,
	}

type Endpoint struct {
	http.Handler
	allow   []string
	methods map[string]http.Handler
}

func NewEndpoint() *Endpoint {
	e := &Endpoint{}
	e.methods = make(map[string]http.Handler)
	e.AddHandler(http.MethodOptions, MiddlewareFunc(http.HandlerFunc(e.OptionsResponse), LogToStderr))
	return e
}

func (e *Endpoint) AddHandler(method string, function http.Handler) error {
	for _, method_check := range HTMLMethods {
		if method == method_check {
			e.methods[method] = function
			e.allow = append(e.allow, method)
			return nil
		}
	}
	return errors.New("no valid method")
}

func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for method, handler := range e.methods {
		if r.Method == method {
			handler.ServeHTTP(w, r)
			return
		}
	}
	ErrorResponse(w, r.Method + " not accepted at this endpoint", http.StatusMethodNotAllowed)
}

func (e *Endpoint) OptionsResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=604800") // 1 week in seconds
	w.Header().Set("Allow-Control-Allow-Origin", "http://www.formicid.io")
	w.Header().Set("Allow", strings.Join(e.allow[:], ", "))
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(e.allow[:], ", "))
}