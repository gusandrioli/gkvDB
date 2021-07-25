package main

func main() {
	db = newDBConnection("")
	defer db.DB.Close()

	transactionStack = NewTransactionStack()

	p := newPrompt()
	p.Run()
}
