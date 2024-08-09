package db

import (
	"database/sql"
	"densho/dict"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var schemaSQL = `
CREATE TABLE IF NOT EXISTS dictionary (
	id INTEGER PRIMARY KEY NOT NULL,
	kanji STRING NOT NULL,
	kana STRING NOT NULL,
	translation STRING NOT NULL
);
`

var insertSQL = `
INSERT into dictionary (
	kanji, kana, translation
) VALUES (
	?, ?, ?
)
`

var listSQL = `
SELECT * FROM dictionary
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

func (db *Db) List() ([]dict.DictEntry, error) {
	tx, err := db.sql.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(listSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []dict.DictEntry

	for rows.Next() {
		var entry dict.DictEntry

		if err := rows.Scan(&entry.Id, &entry.Kanji, &entry.Kana, &entry.Translation); err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return entries, nil
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
