package main

import (
	"fmt"
	"os"
	"strings"
)

func executor(in string) {
	args := strings.Fields(in)
	switch args[0] {
	case "BEGIN":
		Begin(transactionStack, args)
	case "COMMIT":
		Commit(transactionStack, args)
	case "COUNT":
		Count(args)
	case "DELETE":
		Delete(transactionStack, args)
	case "END":
		End(transactionStack, args)
	case "EXIT":
		os.Exit(0)
	case "GET":
		Get(transactionStack, args)
	case "LIST":
		List(args)
	case "NEW":
		NewDB(args)
	case "ROLLBACK":
		Rollback(transactionStack, args)
	case "SET":
		Set(transactionStack, args)
	case "USE":
		UseDB(args)
	default:
		fmt.Printf("ERROR: Unrecognised Command: %s\n", args[0])
	}
}
