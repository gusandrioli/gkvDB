package main

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

func completer(in prompt.Document) []prompt.Suggest {
	if in.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}

	args := strings.Split(in.TextBeforeCursor(), " ")

	return suggestWithArguments(args)

}

func suggestWithArguments(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(GetAllSuggestions(), args[0], true)
	}

	if len(args) >= 3 {
		return []prompt.Suggest{}
	}

	cmd := args[0]

	switch cmd {
	case "COUNT":
		suggestions := []prompt.Suggest{DatabasesCommand, RecordsCommand}
		return prompt.FilterHasPrefix(suggestions, args[1], true)
	case "DELETE":
		suggestions := []prompt.Suggest{DatabaseCommand, RecordCommand}
		return prompt.FilterHasPrefix(suggestions, args[1], true)
	case "LIST":
		suggestions := []prompt.Suggest{DatabasesCommand, RecordsCommand, TransactionsCommand}
		return prompt.FilterHasPrefix(suggestions, args[1], true)
	case "NEW":
		suggestions := []prompt.Suggest{DatabaseCommand}
		return prompt.FilterHasPrefix(suggestions, args[1], true)
	default:
		return []prompt.Suggest{}
	}
}
