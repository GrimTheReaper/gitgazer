package network

import (
	"fmt"
	"net/http"
)

// Serve will serve this application's network stack on provided host and port.
// WARNING: BLOCK
func Serve(host string, port int) (api *API, err error) {
	api = &API{
		host:   host,
		port:   port,
		server: &RegexpHandler{},
	}

	api.registerHandlers()

	return api, api.serve()
}

// API is an abstracted form of our API. Written so you can have more than one!
type API struct {
	host   string
	port   int
	server *RegexpHandler
}

func (api *API) serve() error {
	return http.ListenAndServe(
		fmt.Sprintf("%v:%v", api.host, api.port),
		api.server,
	)
}

func (api *API) registerHandlers() {
	api.server.HandleFunc(`/api/v0/github/user/(\w*)/repositories`, getRepositoriesForUser)
	api.server.HandleFunc(`/api/v0/github/user/(\w*)/followers`, getFollowersForUser)
}
