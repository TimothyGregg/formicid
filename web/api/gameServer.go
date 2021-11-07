package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	ep "github.com/TimothyGregg/formicid/web/api/endpoints"
	"github.com/TimothyGregg/formicid/web/api/storage"
	"github.com/TimothyGregg/formicid/web/util"
	"github.com/gorilla/mux"
)

type GameServer struct {
	http.Server
	store *storage.Store
}

func New_GameServer() *GameServer {
	gs := &GameServer{}
	gs.store = storage.NewStore()

	// build router and middleware stack
	router := mux.NewRouter().StrictSlash(true)

	// middleware routing
	stdMiddleware := []util.Middleware{
		util.LogToStderr,
	}

	getMiddleware := []util.Middleware{
		util.EnforceContentType_JSON,
		util.AddAllowedOrigin,
	}

	// endpoint creation
	homeEP := util.NewEndpoint()
	homeEP.AddHandler(http.MethodGet, ep.HomeHandler(gs.store))

	gameEP := util.NewEndpoint()
	gameEP.AddHandler(http.MethodGet, util.MiddlewareFunc(ep.GameGet(gs.store), getMiddleware...))
	gameEP.AddHandler(http.MethodPost, ep.GamePost(gs.store))

	gameIDEP := util.NewEndpoint()
	gameIDEP.AddHandler(http.MethodGet, util.MiddlewareFunc(ep.ReturnGameByID(gs.store), getMiddleware...))

	endpoints := map[string]*util.Endpoint{
		"/":       homeEP,
		"/g":      gameEP,
		"/g/{id}": gameIDEP,
	}

	// add endpoints to router with middleware
	for endpoint_path, endpoint := range endpoints {
		router.Handle(endpoint_path, util.MiddlewareStack(endpoint, stdMiddleware...))
	}
	// assign
	gs.Handler = router

	// get the appropriate port to serve on
	port := os.Getenv("PORT")
	gs.Addr = ":" + port
	fmt.Println(gs.Addr)

	return gs
}

func (gs *GameServer) Start() {
	log.Fatal(gs.ListenAndServe())
}
