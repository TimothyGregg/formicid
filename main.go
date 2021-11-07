package main

import (
	"sync"

	api "github.com/TimothyGregg/formicid/web/api"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	gs := api.New_GameServer()

	go func() {
		gs.Start()
		wg.Done()
	}()

	wg.Wait()
}