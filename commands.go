package main

import "github.com/c-bata/go-prompt"

var (
	DatabasesCommand = prompt.Suggest{Text: "DATABASES", Description: "Find all databases"}
	DatabaseCommand  = prompt.Suggest{Text: "DATABASE", Description: ""}
	RecordsCommand   = prompt.Suggest{Text: "RECORDS", Description: "Find all records"}
	RecordCommand    = prompt.Suggest{Text: "RECORD", Description: ""}
	DBCommands       = []prompt.Suggest{
		{Text: "NEW", Description: "Creates new database"},
		{Text: "USE", Description: "Use a specific database"},
	}
	TXCommands = []prompt.Suggest{
		{Text: "BEGIN", Description: "Begins a transaction"},
		{Text: "COMMIT", Description: "Commits a transaction"},
		{Text: "COUNT", Description: "Retreives the number of key/velues stored"},
		{Text: "DELETE", Description: "Deletes a value based on a key"},
		{Text: "END", Description: "Ends a transaction"},
		{Text: "EXIT", Description: "Exits the console"},
		{Text: "GET", Description: "Gets a value based on a key"},
		{Text: "LIST", Description: "Lists all databases/Lists all key/values stored"},
		{Text: "ROLLBACK", Description: "Rolls back a transaction"},
		{Text: "SET", Description: "Sets a key to a certain value"},
	}
)

func GetAllSuggestions() []prompt.Suggest {
	cmd := []prompt.Suggest{}

	for _, v := range DBCommands {
		cmd = append(cmd, v)
	}

	for _, v := range TXCommands {
		cmd = append(cmd, v)
	}

	return cmd
}
