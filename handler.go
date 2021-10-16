package main

import (
	"fmt"
)

func Begin(ts *TransactionStack, args []string) {
	ts.Push()
}

func Commit(ts *TransactionStack, args []string) {
	ts.Commit()
	ts.Pop()
}

func Count(args []string) {
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
		CountRecords()
		return
	}

	fmt.Printf(MsgError+"Unrecognized Command: %s\n", args[1])
}

func Delete(ts *TransactionStack, args []string) {
	if len(args) <= 1 {
		fmt.Printf(MsgError + "Missing Command for DELETE")
		return
	}

	if args[1] == "DATABASE" {
		DeleteDatabase(args)
		return
	}

	if args[1] == "RECORD" {
		ts.DeleteRecord(args[2])
		return
	}

	fmt.Printf(MsgError+"Unrecognized Command: %s\n", args[1])
}

func End(ts *TransactionStack, args []string) {
	ts.Pop()
}

func Get(ts *TransactionStack, args []string) {
	ts.Get(args[1])
}

func List(ts *TransactionStack, args []string) {
	// TODO add where = matches contain value?
	if args[1] == "DATABASES" {
		if err := ListDatabases(); err != nil {
			fmt.Printf(MsgError + "No Databases were created yet\n")
		}
		return
	}

	if args[1] == "RECORDS" {
		ListRecords()
		return
	}

	if args[1] == "TRANSACTIONS" {
		ts.ListTransaction()
		return
	}

	fmt.Printf(MsgError+"Unrecognized Command: %s\n", args[1])
}

func NewDB(args []string) {
	var dbName string

	if args[2] != "" {
		dbName = args[2]
	}

	db.DB.Close() // close initial temporary connection
	db = newDBConnection(dbName)
}

func Rollback(ts *TransactionStack, args []string) {
	ts.Rollback()
}

func Set(ts *TransactionStack, args []string) {
	ts.Set(args[1], args[2])
}

func UseDB(args []string) {
	var dbName string

	if args[1] == "" {
		fmt.Println(MsgError + "Missing name\n")
		return
	}

	dbName = args[1]

	db.DB.Close() // close initial temporary connection
	db = newDBConnection(dbName)
}
