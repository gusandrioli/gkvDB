package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gusandrioli/gkvDB/db"
)

func Begin(gkv *gkvDB, args []string) {
	gkv.Push()
}

func Commit(gkv *gkvDB, args []string) {
	gkv.Commit()
	gkv.Pop()
}

func Count(gkv *gkvDB, args []string) {
	if len(args) <= 1 {
		fmt.Printf(MsgError + "Missing Command for COUNT")
		return
	}

	if args[1] == "DATABASES" {
		if err := CountDatabases(); err != nil {
			fmt.Printf(MsgError + "No Databases were created yet\n")
		}
		return
	}

	if args[1] == "RECORDS" {
		gkv.CountRecords()
		return
	}

	fmt.Printf(MsgError+"Unrecognized Command: %s\n", args[1])
}

func Delete(gkv *gkvDB, args []string) {
	if len(args) <= 1 {
		fmt.Printf(MsgError + "Missing Command for DELETE")
		return
	}

	if args[1] == "DATABASE" {
		DeleteDatabase(args)
		return
	}

	if args[1] == "RECORD" {
		gkv.DeleteRecord(args[2])
		return
	}

	fmt.Printf(MsgError+"Unrecognized Command: %s\n", args[1])
}

func End(gkv *gkvDB, args []string) {
	gkv.Pop()
}

// TODO
func Expire(gkv *gkvDB, args []string) {
	min, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf(MsgError + "Invalid time format.")
	}

	expire := time.Now().Add(time.Duration(min) * time.Minute)
	gkv.SetExpire(args[1], expire.String())
}

func Get(gkv *gkvDB, args []string) {
	gkv.Get(args[1])
}

func List(gkv *gkvDB, args []string) {
	if args[1] == "DATABASES" {
		if err := ListDatabases(); err != nil {
			fmt.Printf(MsgError + "No Databases were created yet\n")
		}
		return
	}

	if args[1] == "RECORDS" {
		gkv.ListRecords()
		return
	}

	if args[1] == "TRANSACTIONS" {
		gkv.ListTransaction()
		return
	}

	fmt.Printf(MsgError+"Unrecognized Command: %s\n", args[1])
}

func NewDB(gkv *gkvDB, args []string) {
	var dbName string

	if args[2] != "" {
		dbName = args[2]
	}

	gkv.mainDB.DB.Close() // close initial temporary connection
	gkv.mainDB = db.NewDBConnection(dbName)
}

func Rollback(gkv *gkvDB, args []string) {
	gkv.Rollback()
}

func Set(gkv *gkvDB, args []string) {
	gkv.Set(args[1], args[2])
}

func UseDB(gkv *gkvDB, args []string) {
	var dbName string

	if args[1] == "" {
		fmt.Println(MsgError + "Missing name\n")
		return
	}

	dbName = args[1]

	gkv.mainDB.DB.Close() // close initial temporary connection
	gkv.mainDB = db.NewDBConnection(dbName)
}
