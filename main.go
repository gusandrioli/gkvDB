package main

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

var GlobalStore = make(map[string]string)

type Transaction struct {
	localStore map[string]string
	next       *Transaction // linked list so it points to the next Transaction
}

type TransactionStack struct {
	topTransaction *Transaction
	size           int
}

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

func (ts *TransactionStack) Commit() {
	activeTransaction := ts.GetTopTransaction()

	if activeTransaction == nil {
		fmt.Printf("WARNING: Nothing to commit")
		return
	}

	for k, v := range activeTransaction.localStore {
		GlobalStore[k] = v

		if activeTransaction.next != nil {
			activeTransaction.next.localStore[k] = v
		}
	}
}

func Count() {
	fmt.Printf("%d\n", len(GlobalStore))
}

func Delete(key string, ts *TransactionStack) {
	activeTransaction := ts.GetTopTransaction()

	if activeTransaction != nil {
		if _, ok := activeTransaction.localStore[key]; ok {
			delete(activeTransaction.localStore, key)
			return
		}
	}

	delete(GlobalStore, key)
}

func Get(key string, ts *TransactionStack) {
	activeTransaction := ts.GetTopTransaction()

	if activeTransaction != nil {
		if v, ok := activeTransaction.localStore[key]; ok {
			fmt.Printf("%s\n", v)
			return
		}
	}

	if v, ok := GlobalStore[key]; ok {
		fmt.Printf("%s\n", v)
		return
	}

	fmt.Printf("%s not set\n", key)
	return
}

func (ts *TransactionStack) GetTopTransaction() *Transaction {
	return ts.topTransaction
}

func List() {
	if GlobalStore == nil {
		return
	}

	for k, v := range GlobalStore {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func (ts *TransactionStack) Push() {
	pushedTransaction := &Transaction{
		localStore: make(map[string]string),
	}

	pushedTransaction.next = ts.topTransaction
	ts.topTransaction = pushedTransaction
	ts.size++
}

func (ts *TransactionStack) Pop() {
	if ts.topTransaction == nil {
		fmt.Printf("ERROR: No Active Transactions\n")
		return
	}

	ts.topTransaction = ts.topTransaction.next
	ts.size--
}

func (ts *TransactionStack) Rollback() {
	if ts.topTransaction == nil {
		fmt.Printf("ERROR: No Active Transactions")
		return
	}

	for k := range ts.topTransaction.localStore {
		delete(ts.topTransaction.localStore, k)
	}
}

func Set(key string, value string, ts *TransactionStack) {
	activeTransaction := ts.GetTopTransaction()

	if activeTransaction != nil {
		activeTransaction.localStore[key] = value
		return
	}

	GlobalStore[key] = value
	return
}

func executor(in string) {
	fmt.Println("Your input is here: " + in)
	if in == "" {
		LivePrefixState.IsEnable = false
		LivePrefixState.LivePrefix = in
		return
	}

	LivePrefixState.LivePrefix = in + "> "
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
		fmt.Printf("ERROR: Unrecognised Command %s\n", commands[0])
	}

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

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("gkvDB >>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("gkvDB"),
	)
	p.Run()
}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Printf("gkv > ")
// 		text, _ := reader.ReadString('\n')

// 	}
