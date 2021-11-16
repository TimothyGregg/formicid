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
		util.FixCORS,
	}

	// endpoint creation
	homeEP := util.NewEndpoint()
	homeEP.AddHandler(http.MethodGet, ep.HomeHandler(gs.store))

	gamesEP := util.NewEndpoint()
	gamesEP.AddHandler(http.MethodGet, ep.GameGet(gs.store))
	//gamesEP.AddHandler(http.MethodPost, ep.GamePost(gs.store))

	gameIDEP := util.NewEndpoint()
	gameIDEP.AddHandler(http.MethodGet, ep.ReturnGameByID(gs.store))

	endpoints := map[string]*util.Endpoint{
		"/":           homeEP,
		"/games":      gamesEP,
		"/games/{id}": gameIDEP,
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
	oc := os.Getenv("FORMICIDIOOC")
	if oc == "" {
		fmt.Println("No Cert: \"" + oc + "\"")
	}
	pk := os.Getenv("FORMICIDIOPK")
	if pk == "" {
		fmt.Println("No Key: \"" + pk + "\"")
	}
	fmt.Println(os.Environ())
	log.Fatal(gs.ListenAndServeTLS(oc, pk))
}
