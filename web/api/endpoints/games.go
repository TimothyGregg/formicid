package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TimothyGregg/formicid/web/api/storage"
	"github.com/TimothyGregg/formicid/web/util"
)

func GameGet(s *storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Game get!")
		if len(s.Games) == 0 {
			util.Response_NotFound(w, "No Games at all *shrug*")
		} else {
			bytes, err := json.MarshalIndent(g.games, "", "\t")
		if err != nil {
			util.Res
		}	
	}
}