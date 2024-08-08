package api

import (
	"densho/dict"
	"encoding/json"
	"net/http"
)

type DictController struct {
	dictionary *dict.Dict
}

func NewDictController(dictionary *dict.Dict) DictController {
	return DictController{dictionary: dictionary}
}

func (dc *DictController) GetEntries(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(dc.dictionary.Entries())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	}

	dc.dictionary.AddEntry(&entry)

	w.WriteHeader(http.StatusCreated)
}

func (dc *DictController) UpdateEntry(w http.ResponseWriter, r *http.Request) {
	var updatedEntry dict.DictEntry

	err := json.NewDecoder(r.Body).Decode(&updatedEntry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	dc.dictionary.UpdateEntry(&updatedEntry)

	w.WriteHeader(http.StatusOK)
}
