package frontend

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FrontendServer struct {
	http.Server
}


func New_FrontendServer() (*FrontendServer) {
	fs := &FrontendServer{}

	// build router
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))
	fs.Handler = router

	return fs
}

func (fs *FrontendServer) Start() {
	log.Fatal(fs.ListenAndServe())
}