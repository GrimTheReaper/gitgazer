package network

import (
	"encoding/json"
	"net/http"
	"regexp"
)

/*
Source: https://stackoverflow.com/a/6565407
Although I modified it a bit.
*/

type route struct {
	pattern *regexp.Regexp
	handler func(http.ResponseWriter, *http.Request, *regexp.Regexp)
}

// RegexpHandler was written by "Evan Shaw" over on Stackoverflow
type RegexpHandler struct {
	routes []*route
}

// HandleFunc will append the handler function to the routes
func (h *RegexpHandler) HandleFunc(regex string, handler func(http.ResponseWriter, *http.Request, *regexp.Regexp)) {
	h.routes = append(h.routes, &route{regexp.MustCompile(regex), handler})
}

// ServeHTTP is for http.ListenAndServe
func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler(w, r, route.pattern)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

func getFirstCaptureGroup(request *http.Request, pattern *regexp.Regexp) string {
	return pattern.FindStringSubmatch(request.URL.Path)[1] // We can assume that it will be there because it was vetted earlier in the code.
}

func serveFormatted(rw http.ResponseWriter, object interface{}) {
	byts, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		rw.Write([]byte("Failed to encode response"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Write(byts)
}
