package endpoints

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TimothyGregg/formicid/web/api/storage"
	"github.com/TimothyGregg/formicid/web/util"
	"github.com/gorilla/mux"
)

func ReturnGameByID(s *storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key, _ := strconv.ParseInt(vars["id"], 10, 64)
		if int(key) < len(s.Games) {
			json.NewEncoder(w).Encode(s.Games[int(key)])
		} else {
			util.Response_NotFound(w, "No game with that ID exists")
		}
	}
}