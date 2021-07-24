package main

import "fmt"

var GlobalStore = make(map[string]string)

type Transaction struct {
	localStore map[string]string
	next       *Transaction // linked list so it points to the next Transaction
}

type TransactionStack struct {
	topTransaction *Transaction
	size           int
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
