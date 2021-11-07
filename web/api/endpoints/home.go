package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TimothyGregg/formicid/web/api/storage"
	"github.com/TimothyGregg/formicid/web/util"
)

func HomeHandler(s *storage.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Home handler!")
		_, err := ioutil.ReadAll(r.Body) // body, err
		if err != nil {
			util.Response_BadRequest(w)
			return
		}
		json.NewEncoder(w).Encode("You found the home page!")
	}
}