package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/TimothyGregg/Antmound/game"
	"github.com/TimothyGregg/Antmound/game/tools"
	"github.com/gorilla/mux"
)

type Server struct {
	Games []*game.Game
	UID_Generator *tools.UID_Generator
}

func New_Server() *Server {
	s := &Server{}
	s.UID_Generator = tools.New_UID_Generator()
	return s
}

func (s *Server) New_Game(size_x, size_y, tries int) {
	s.Games = append(s.Games, game.New_Game(s.UID_Generator.Next(), size_x, size_y, tries))
}

func (s *Server) homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func (s *Server) returnGames(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnGames")
	json.NewEncoder(w).Encode(s.Games)
}

func (s *Server) returnGameByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.ParseInt(vars["id"], 10, 64)
	if int(key) < len(s.Games) {
		json.NewEncoder(w).Encode(s.Games[int(key)])
	} else {
		http.Error(w, "Game not found", 404)
	}
}

func (s *Server) HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", s.homePage)
    router.HandleFunc("/games", s.returnGames).Methods("GET")
    router.HandleFunc("/games/{id}", s.returnGameByID).Methods("GET")
	port := os.Getenv("PORT")
	port = ":" + port
    log.Fatal(http.ListenAndServe(port, router))
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