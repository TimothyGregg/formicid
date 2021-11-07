package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TimothyGregg/formicid/web/api/actions"
	"github.com/TimothyGregg/formicid/web/api/storage"
	"github.com/TimothyGregg/formicid/web/util"
)

func GameGet(s *storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Game get!")
		if len(s.Games) == 0 {
			util.Response_NotFound(w, "No Games at all *shrug*")
		} else {
			json.NewEncoder(w).Encode(s.Games)
		}
	}
}

func GamePost(s *storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get response body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			util.Response_BadRequest(w)
			return
		}
		// ensure the action is properly formed
		a, err := actions.NewAction(body)
		if err != nil {
			util.Response_BadRequest(w, "Malformed action")
			return
		}
		fmt.Println("WE GOT TO HERE")
		// perform appropriate action
		switch a.Action_type {
		case actions.META:
			switch a.Action_name {
			case "addGame":
				details := &struct{
					Size_x int `json:"size_x"`
					Size_y int `json:"size_y"`}{}
				json.Unmarshal(a.Action_details, details)
				fmt.Println("Was this the problem?")
				//s.New_Game(details.Size_x, details.Size_y)
			}
		}
	}
}