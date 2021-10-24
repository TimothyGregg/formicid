package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TimothyGregg/formicid/game"
	"github.com/TimothyGregg/formicid/game/tools"
	"github.com/TimothyGregg/formicid/web/util"
	"github.com/gorilla/mux"
)

type GameServer struct {
	http.Server
	Games         []*game.Game
	UID_Generator *tools.UID_Generator
}

func New_GameServer() *GameServer {
	gs := &GameServer{}
	gs.UID_Generator = tools.New_UID_Generator()

	// build router and middleware stack
	router := mux.NewRouter().StrictSlash(true)

	// endpoint creation
	homeEP := &util.Endpoint{
		Default: gs.homeHandler,
	}

	gameEP := &util.Endpoint{
		Get: gs.gameGet,
	}

	gameIDEP := &util.Endpoint{
		Get: gs.returnGameByID,
	}

	endpoints := map[string]*util.Endpoint{
		"/":       homeEP,
		"/g":      gameEP,
		"/g/{id}": gameIDEP,
	}

	// Add endpoints to router
	for endpoint_path, endpoint := range endpoints {
		router.Handle(endpoint_path, endpoint)
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

func (gs *GameServer) New_Game(size_x, size_y int) {
	gs.Games = append(gs.Games, game.New_Game(gs.UID_Generator.Next(), size_x, size_y))
}
