package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/TimothyGregg/formicid/web/api/v1.0/actions"
	"github.com/TimothyGregg/formicid/web/util"
	"github.com/gorilla/mux"
)

func (gs *GameServer) homePost(w http.ResponseWriter, r *http.Request) {

}

func (gs *GameServer) gamePost(w http.ResponseWriter, r *http.Request) {
	// get response body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ErrorResponse(w, "Bad request", http.StatusBadRequest)
		return
	}
	// ensure the action is properly formed
	a, err := actions.NewAction(body)
	if err != nil {
		util.ErrorResponse(w, "Malformed action", http.StatusBadRequest)
	}
	// perform appropriate action
	switch a.Action_type {
	case actions.META:
		switch a.Action_name {
		case "addGame":
			details := struct{
				Size_x int `json:"size_x"`
				Size_y int `json:"size_y"`}{}
			json.Unmarshal(a.Action_details, details)
			gs.New_Game(details.Size_x, details.Size_y)
		}
	}
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
		util.ErrorResponse(w, "Bad request", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("You found the home page!")
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
