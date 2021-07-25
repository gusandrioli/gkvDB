package main

import (
	prompt "github.com/c-bata/go-prompt"
)

func newPrompt() *prompt.Prompt {
	return prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("gkvDB >>> "),
		prompt.OptionTitle("gkvDB"),
	)
}
