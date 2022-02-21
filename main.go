package main

import "github.com/gusandrioli/gkvDB/db"

func main() {
	mainDB := db.NewDBConnection("")
	defer mainDB.DB.Close()
	helperDB := db.NewDBConnection("local_expiry")
	defer helperDB.DB.Close()

	gkv = NewGkv(mainDB, helperDB)
	p := newPrompt()
	p.Run()
}
