package api

import (
	"densho/db"
	"log"
	"net/http"
	"time"
)

type Api struct {
	server *http.Server
}

func NewApi(address string, db *db.Db) Api {
	dictService := NewDictService(db)

	dictionaryController := NewDictController(&dictService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /entries", dictionaryController.GetEntries)
	mux.HandleFunc("POST /entries", dictionaryController.PostEntry)
	mux.HandleFunc("GET /entriess", dictionaryController.FindEntries)

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
