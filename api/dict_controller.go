package api

import (
	"densho/dict"
	"encoding/json"
	"net/http"
)

type DictController struct {
	dictService *DictService
}

func NewDictController(dictionary *DictService) DictController {
	return DictController{dictService: dictionary}
}

func (dc *DictController) GetEntries(w http.ResponseWriter, r *http.Request) {
	entries, err := dc.dictService.Entries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(entries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (dc *DictController) PostEntry(w http.ResponseWriter, r *http.Request) {
	var entry dict.DictEntryPayload

	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dc.dictService.AddEntry(&entry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
