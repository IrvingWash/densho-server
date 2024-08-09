package db

import (
	"database/sql"
	"densho/dict"
	"fmt"
)

var schemaSQL = `
CREATE TABLE IF NOT EXISTS dictionary (
	id NUMBER PRIMARY KEY
	kanji STRING
	kana STRING
	translation STRING
);
`

var insertSQL = `
INSERT into dictionary (
	kanji, kana, translation
) VALUES (
	?, ?, ?
)
`

type Db struct {
	sql       *sql.DB
	insertion *sql.Stmt
}

func NewDb(dbFile string) Db {
	sqlDb, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to open database, %s", err.Error()))
	}

	_, err = sqlDb.Exec(schemaSQL)
	if err != nil {
		panic(fmt.Sprintf("Failed to create table, %s", err.Error()))
	}

	insertion, err := sqlDb.Prepare(insertSQL)
	if err != nil {
		panic(fmt.Sprintf("Failed to prepare insertion, %s", err.Error()))
	}

	return Db{
		sql:       sqlDb,
		insertion: insertion,
	}
}

func (db *Db) Insert(entry dict.DictEntryPayload) error {
	tx, err := db.sql.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Stmt(db.insertion).Exec(entry.Kanji, entry.Kana, entry.Translation)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (db *Db) Close() {
	db.insertion.Close()
	db.sql.Close()
}
