package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gusandrioli/gkvDB/db"
)

func main() {
	l := log.Default()
	helper := db.NewDBConnection("local_expiry")
	iter := helper.DB.NewIterator(nil, nil)

	fmt.Println("Starting cleaner...")

	dbsToClean := make(map[string][]string)
	for iter.Next() {
		fmt.Println(iter.Value())
		expiryDate, _ := time.Parse(time.RFC3339, string(iter.Value()))
		if time.Now().After(expiryDate) {
			keyParts := strings.Split(string(iter.Key()), ".")
			expiredKeys := dbsToClean[keyParts[0]]
			expiredKeys = append(expiredKeys, strings.Join(keyParts[1:], "."))
			dbsToClean[keyParts[0]] = expiredKeys

			fmt.Printf("Removing: %v\n", iter.Key())
			if err := helper.DB.Delete(iter.Key(), nil); err != nil {
				l.Println(err)
			}

		}
	}
	iter.Release()

	if len(dbsToClean) < 1 {
		fmt.Printf("Nothing to clean\n")
		os.Exit(1)
	}

	for dbToClean, expiredKeys := range dbsToClean {
		mainDB := db.NewDBConnection(dbToClean)

		for _, key := range expiredKeys {
			if err := mainDB.DB.Delete([]byte(key), nil); err != nil {
				l.Println(err)
			}
			fmt.Printf("Removed: %v\n", key)
		}

	}
}
