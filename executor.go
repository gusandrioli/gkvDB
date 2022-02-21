package main

import (
	"fmt"
	"os"
	"strings"
)

func executor(in string) {
	args := strings.Fields(in)

	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "BEGIN":
		Begin(gkv, args)
	case "COMMIT":
		Commit(gkv, args)
	case "COUNT":
		Count(gkv, args)
	case "DELETE":
		Delete(gkv, args)
	case "END":
		End(gkv, args)
	case "EXIT":
		os.Exit(0)
	case "EXPIRE":
		Expire(gkv, args)
	case "GET":
		Get(gkv, args)
	case "LIST":
		List(gkv, args)
	case "NEW":
		NewDB(gkv, args)
	case "ROLLBACK":
		Rollback(gkv, args)
	case "SET":
		Set(gkv, args)
	case "USE":
		UseDB(gkv, args)
	default:
		fmt.Printf("ERROR: Unrecognised Command: %s\n", args[0])
	}
}
