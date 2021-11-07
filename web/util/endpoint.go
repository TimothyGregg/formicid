package util

import (
	"errors"
	"net/http"
)

var	HTMLMethods = [6]string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodOptions,
	}

type Endpoint struct {
	http.Handler
	allow   []string
	methods map[string]http.Handler
}

func NewEndpoint() *Endpoint {
	e := &Endpoint{}
	e.methods = make(map[string]http.Handler)
	return e
}

// AddHandler adds a map entry in e.methods to associate a http.Handler to a request method
func (e *Endpoint) AddHandler(method string, methodHandler http.Handler) error {
	for _, method_check := range HTMLMethods {
		if method == method_check {
			e.methods[method] = methodHandler
			e.allow = append(e.allow, method)
			return nil
		}
	}
	return errors.New("invalid method for assignment: " + method)
}

func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Select appropriate method, if supported
	methodHandler, exists := e.methods[r.Method]
	if exists {
		methodHandler.ServeHTTP(w, r)
		return
	}

	// Deal with methods not supported at the endpoint with a "Method Not Allowed" respose
	Response_MethodNotAllowed(w, e.allow)
}

// func (e *Endpoint) OptionsResponse(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Cache-Control", "max-age=604800") // 1 week in seconds
// 	w.Header().Set("Allow", strings.Join(e.allow[:], ", "))
// 	w.Header().Set("Access-Control-Allow-Origin", "http://www.formicid.io")
// 	w.Header().Set("Access-Control-Allow-Methods", strings.Join(e.allow[:], ", "))
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	w.WriteHeader(http.StatusOK)
// }