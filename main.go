package main

import (
	"github.com/TimothyGregg/Antmound/api"
)

func main() {
	s := api.New_Server()
	s.New_Game(100, 100, 10)
	s.HandleRequests()
}