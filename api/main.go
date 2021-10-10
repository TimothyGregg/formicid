package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TimothyGregg/Antmound/game"
)

var Game *game.Game

func Server() {
	Game = game.New_Game(100, 100, 10)
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnBoard")
	json.NewEncoder(w).Encode(Game.Board)
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/board", returnBoard)
    log.Fatal(http.ListenAndServe(":10000", nil))
}


/*
func Webpage() {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)
	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
*/