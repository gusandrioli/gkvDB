package db

import (
	"fmt"
	"log"
	"strings"

	"github.com/gusandrioli/gkvDB/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDB struct {
	Name string
	DB   *leveldb.DB
}

func NewDBConnection(name string) *LevelDB {
	var dbName string

	if name == "" {
		dbName = "tmp-" + utils.NewSequenceWithLength(8)
	} else {
		dbName = name
	}

	db, err := leveldb.OpenFile("./.tmp/"+dbName+".db", nil)
	if err != nil {
		log.Fatal(err, "leveldb.Openfile")
		return nil
	}

	if !strings.Contains(dbName, "tmp-") && !strings.Contains(dbName, "local_expiry") {
		fmt.Printf("Database initialized: %s\n", dbName)
	}

	return &LevelDB{
		Name: dbName,
		DB:   db,
	}
}
