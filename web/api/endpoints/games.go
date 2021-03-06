package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/TimothyGregg/formicid/db"
	"github.com/TimothyGregg/formicid/game"
	"github.com/TimothyGregg/formicid/web/api/actions"
	"github.com/TimothyGregg/formicid/web/api/storage"
	"github.com/TimothyGregg/formicid/web/util"
)

func GameGet(s *storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(s.Games) == 0 {
			util.Response_NotFound(w, "No Games at all *shrug*")
			return
		} else {
			bytes, err := json.MarshalIndent(s.Games, "", "\t")
			if err != nil {
				util.Response_ServerUnavailable(w)
				return
			}
			w.Write(bytes)
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
		}
		// perform appropriate action
		database, err := db.GetDatabase()
		if err != nil {
			util.Response_ServerUnavailable(w, err.Error())
		}
		g := game.New_Game(s.UID_Generator.Next(), 100, 100)
		err = database.StoreGame("/", g)
		if err != nil {
			util.Response_ServerUnavailable(w, err.Error())
		}
	}
}
