package api

import (
	"densho/db"
	"densho/dict"
)

type DictService struct {
	database *db.Db
}

func NewDictService(database *db.Db) DictService {
	return DictService{database: database}
}

func (ds *DictService) Entries() ([]dict.DictEntry, error) {
	return ds.database.List()
}

func (ds *DictService) FindEntries(query string) ([]dict.DictEntry, error) {
	return ds.database.Find(query)
}

func (ds *DictService) AddEntry(entry *dict.DictEntryPayload) error {
	return ds.database.Insert(*entry)
}
