package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/TimothyGregg/formicid/game"
	"github.com/TimothyGregg/formicid/game/tools"
	"github.com/gorilla/mux"
)

type GameServer struct {
	http.Server
	Games []*game.Game
	UID_Generator *tools.UID_Generator
}

func New_GameServer() (*GameServer) {
	gs := &GameServer{}
	gs.UID_Generator = tools.New_UID_Generator()

	// build router
	router := mux.NewRouter().StrictSlash(true)	
	router.HandleFunc("/", gs.homePost)
	router.HandleFunc("/games", gs.returnGames).Methods("GET")
    router.HandleFunc("/games/{id}", gs.returnGameByID).Methods("GET")
	gs.Handler = router

	// get the appropriate port to serve on
	port := os.Getenv("PORT")
	gs.Addr = ":" + port

	return gs
}

func (gs *GameServer) Start() {
	log.Fatal(gs.ListenAndServe())
}

func (gs *GameServer) New_Game(size_x, size_y int) {
	gs.Games = append(gs.Games, game.New_Game(gs.UID_Generator.Next(), size_x, size_y))
}

func (gs *GameServer) homePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, "Bad action", 400)
	}
	headerContentType := r.Header.Get("Content-Type") //https://golangbyexample.com/validate-range-http-body-golang/
	if headerContentType != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	
}

func (gs *GameServer) returnGames(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnGames")
	if len(gs.Games) == 0 {
		errorResponse(w, "No Games", http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(gs.Games)
	}
}

func (gs *GameServer) returnGameByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.ParseInt(vars["id"], 10, 64)
	if int(key) < len(gs.Games) {
		json.NewEncoder(w).Encode(gs.Games[int(key)])
	} else {
		errorResponse(w, "Game not found", http.StatusNotFound)
	}
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func ok(w http.ResponseWriter) {
	errorResponse(w, "Success", http.StatusOK)
}
