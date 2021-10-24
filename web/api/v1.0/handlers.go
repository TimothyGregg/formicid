package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/TimothyGregg/formicid/web/util"
	"github.com/gorilla/mux"
)

func (gs *GameServer) homePost(w http.ResponseWriter, r *http.Request) {

}

func (gs *GameServer) gameGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Game get!")
	if len(gs.Games) == 0 {
		util.ErrorResponse(w, "No Games", http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(gs.Games)
	}
}

func (gs *GameServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home handler!")
	_, err := ioutil.ReadAll(r.Body) // body, err
	if err != nil {
		util.ErrorResponse(w, "Bad action", 400)
	}
	json.NewEncoder(w).Encode("You found the home page!")
	util.ErrorResponse(w, "Shit", http.StatusTeapot)
}

func (gs *GameServer) returnGameByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Game ID get!")
	vars := mux.Vars(r)
	key, _ := strconv.ParseInt(vars["id"], 10, 64)
	if int(key) < len(gs.Games) {
		json.NewEncoder(w).Encode(gs.Games[int(key)])
	} else {
		util.ErrorResponse(w, "Game not found", http.StatusNotFound)
	}
}
