package main

import (
	"fmt"
	"sync"

	"github.com/TimothyGregg/Antmound/api"
	"github.com/TimothyGregg/Antmound/frontend"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	gs := api.New_GameServer()
	fs := frontend.New_FrontendServer()

	go func() {
		gs.Start()
		wg.Done()
	}()
	fmt.Println("1")
	go func() {
		fs.Start()
		wg.Done()
	}()
	fmt.Println("2")

	wg.Wait()
}