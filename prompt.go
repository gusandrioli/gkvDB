package main

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "BEGIN", Description: "Begins a transaction"},
		{Text: "COMMIT", Description: "Commits a transaction."},
		{Text: "COUNT", Description: "Retreives the number of key/velues stored"},
		{Text: "DELETE", Description: "Deletes a value based on a key"},
		{Text: "END", Description: "Ends a transaction"},
		{Text: "EXIT", Description: "Exits the console"},
		{Text: "GET", Description: "Gets a value based on a key"},
		{Text: "LIST", Description: "Lists all key/values stored"},
		{Text: "ROLLBACK", Description: "Rolls back a transaction"},
		{Text: "SET", Description: "Sets a key to a certain value"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func executor(in string) {
	if in == "" {
		LivePrefixState.IsEnable = false
		return
	}

	LivePrefixState.IsEnable = true

	ts := &TransactionStack{}
	commands := strings.Fields(in)
	switch commands[0] {
	case "BEGIN":
		ts.Push()
	case "ROLLBACK":
		ts.Rollback()
	case "COMMIT":
		ts.Commit()
		ts.Pop()
	case "END":
		ts.Pop()
	case "SET":
		Set(commands[1], commands[2], ts)
	case "GET":
		Get(commands[1], ts)
	case "LIST":
		List()
	case "DELETE":
		Delete(commands[1], ts)
	case "COUNT":
		Count()
	case "EXIT":
		os.Exit(0)
	default:
		fmt.Printf("ERROR: Unrecognised Command: %s\n", commands[0])
	}
}

func newPrompt() *prompt.Prompt {
	return prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("gkvDB >>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("gkvDB"),
	)
}
