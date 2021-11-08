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
		_, err = actions.NewAction(body)
		if err != nil {
			util.Response_BadRequest(w, "Malformed action")
			return
		}
		fmt.Println("WE GOT TO HERE")

		s.New_Game(500, 500)
	}
}