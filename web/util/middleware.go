package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// https://medium.com/@chrisgregory_83433/chaining-middleware-in-go-918cfbc5644d
// https://www.alexedwards.net/blog/making-and-using-middleware
type Middleware func(http.Handler) http.Handler

func MiddlewareStack(finalHandler http.Handler, m ...Middleware) http.Handler {
	if len(m) < 1 {
		return finalHandler
	}
	wrapped := finalHandler

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}
	return wrapped
}

func MiddlewareFunc(finalHandlerFunc http.HandlerFunc, m ...Middleware) http.Handler {
	return MiddlewareStack(finalHandlerFunc, m...)
}

func EnforceContentType_JSON(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerContentType := r.Header.Get("Content-Type") //https://golangbyexample.com/validate-range-http-body-golang/
		if headerContentType != "application/json" {
			Response_UnsupportedMediaType(w, "Content Type is not application/json")
			return
		}
		nextHandler.ServeHTTP(w, r)
	})
}

// https://blog.questionable.services/article/guide-logging-middleware-go/
// https://stackoverflow.com/questions/40396499/go-create-io-writer-inteface-for-logging-to-mongodb-database
func LogToStderr(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Response_BadRequest(w)
		}
		header := ""
		for k, v := range r.Header {
			header = header + k + ":"
			for _, s := range v {
				header = header + s + ","
			}
			header = header + "\n"
		}
		fmt.Fprintf(os.Stderr, "Method:%s\nHeaders: %s\nBody:%s\n", r.Method, header, body)

		nextHandler.ServeHTTP(w, r)
	})
}

// Access-Control-Allow-Origin
func FixCORS(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://www.formicid.io")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin")
		nextHandler.ServeHTTP(w, r)
	})
}