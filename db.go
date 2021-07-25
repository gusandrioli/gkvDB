package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	db *LevelDB
)

type LevelDB struct {
	DB *leveldb.DB
}

func newDBConnection(name string) *LevelDB {
	var dbName string

	if name == "" {
		dbName = "tmp-" + newSequenceWithLength(8)
	} else {
		dbName = name
	}

	db, err := leveldb.OpenFile("./.tmp/"+dbName+".db", nil)
	if err != nil {
		log.Fatal(err, "leveldb.Openfile")
		return nil
	}

	if !strings.Contains(dbName, "tmp-") {
		fmt.Printf("Database initialized: %s\n", dbName)
	}

	return &LevelDB{
		DB: db,
	}
}
