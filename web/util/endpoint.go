package util

import (
	"errors"
	"net/http"
	"strings"
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
	e.AddHandler(http.MethodOptions, OptionsResponse(e))
	return e
}

// AddHandler adds a map entry in e.methods to associate a http.Handler to a request method
func (e *Endpoint) AddHandler(method string, methodHandler http.Handler) error {
	for _, method_check := range HTMLMethods {
		if method == method_check {
			e.allow = append(e.allow, method)
			e.methods[method] = methodHandler
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

func OptionsResponse(e *Endpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerContent := make(map[string]string)
		headerContent["Cache-Control"] = "max-age=604800" // 1 week in seconds
		headerContent["Allow"] = strings.Join(e.allow[:], ", ")
		headerContent["Access-Control-Allow-Origin"] = "http://www.formicid.io"
		headerContent["Access-Control-Allow-Methods"] = strings.Join(e.allow[:], ", ")
		headerContent["Access-Control-Allow-Headers"] = "Content-Type, Access-Control-Allow-Origin"
		Response_NoContent(w, headerContent)
	}
}