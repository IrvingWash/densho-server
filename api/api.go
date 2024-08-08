package api

import (
	"densho/dict"
	"log"
	"net/http"
	"time"
)

type Api struct {
	server *http.Server
}

func NewApi(address string) Api {
	dictionary := dict.NewDict()

	dictionaryController := NewDictController(&dictionary)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /entries", dictionaryController.GetEntries)
	mux.HandleFunc("POST /entries", dictionaryController.PostEntry)

	s := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return Api{server: s}
}

func (api *Api) Start() {
	log.Fatal(api.server.ListenAndServe())
}
