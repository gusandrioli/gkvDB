package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

var (
	transactionStack = &TransactionStack{}
)

type Transaction struct {
	localStore map[string]string
	next       *Transaction
}

type TransactionStack struct {
	topTransaction *Transaction
	size           int
}

func (ts *TransactionStack) Commit() {
	activeTransaction := ts.GetTopTransaction()

	if activeTransaction == nil {
		fmt.Printf(MsgWarning + "Nothing to commit")
		return
	}

	for k, v := range activeTransaction.localStore {
		if err := db.DB.Put([]byte(k), []byte(v), nil); err != nil {
			fmt.Printf(MsgError + err.Error() + "\n")
		}

		if activeTransaction.next != nil {
			activeTransaction.next.localStore[k] = v
		}
	}
}

func CountRecords() {
	var i int64
	iter := db.DB.NewIterator(nil, nil)
	for iter.Next() {
		i++
	}
	iter.Release()

	fmt.Printf("%d record(s) were found\n", i)
}

func CountDatabases() error {
	dbs, err := ioutil.ReadDir("./.tmp/")
	if err != nil {
		return errors.Wrap(err, "ioutil.ReadDir")
	}

	fmt.Printf("%d Databases found\n", len(dbs))

	return nil
}

func (ts *TransactionStack) DeleteRecord(key string) {
	activeTransaction := ts.GetTopTransaction()

	if activeTransaction != nil {
		if _, ok := activeTransaction.localStore[key]; ok {
			delete(activeTransaction.localStore, key)
			return
		}
	}

	if err := db.DB.Delete([]byte(key), nil); err != nil {
		fmt.Printf(MsgError + err.Error() + "\n")
	}
}

func DeleteDatabase(args []string) {
	if len(args) <= 2 {
		fmt.Printf(MsgError + "Missing database name for DELETE\n")
		return
	}

	os.RemoveAll("./.tmp/" + args[2] + ".db")
}

func (ts *TransactionStack) Get(key string) {
	activeTransaction := ts.GetTopTransaction()

	if activeTransaction != nil {
		if v, ok := activeTransaction.localStore[key]; ok {
			fmt.Printf("%s\n", v)
			return
		}
	}

	value, err := db.DB.Get([]byte(key), nil)
	if err != nil {
		fmt.Printf(MsgError+"%s not found\n", key)
		return
	}

	fmt.Printf("%s\n", string(value))
}

func (ts *TransactionStack) GetTopTransaction() *Transaction {
	return ts.topTransaction
}

func ListDatabases() error {
	dbs, err := ioutil.ReadDir("./.tmp/")
	if err != nil {
		return errors.Wrap(err, "ioutil.ReadDir")
	}

	fmt.Printf("%d Databases found: \n", len(dbs))

	for i, v := range dbs {
		name := []byte(v.Name())
		fmt.Printf("%d: %s\n", i, string(name[:len(name)-3]))
	}

	return nil
}

func ListRecords() {
	iter := db.DB.NewIterator(nil, nil)
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		fmt.Printf("%s: %s\n", k, v)
	}
	iter.Release()

	if err := iter.Error(); err != nil {
		fmt.Printf(MsgError + err.Error() + "\n")
	}
}

func (ts *TransactionStack) ListTransaction() {
	if ts.topTransaction == nil {
		fmt.Printf("No changes in transaction\n")
		return
	}

	transactionStackChanges := []map[string]string{}

	currentTransaction := ts.topTransaction

	for {
		transactionStackChanges = append(transactionStackChanges, currentTransaction.localStore)
		if currentTransaction.next == nil {
			break
		}

		currentTransaction = currentTransaction.next
	}

	for transactionIndex := range transactionStackChanges {
		invTransactionIndex := len(transactionStackChanges) - 1 - transactionIndex // hacky way to invert count in loop

		for k, v := range transactionStackChanges[transactionIndex] {
			fmt.Printf("Transaction Nesting Level: %d => %s: %s\n", invTransactionIndex, k, v)
		}
	}
}

func NewTransactionStack() *TransactionStack {
	return &TransactionStack{}
}

func (ts *TransactionStack) Push() {
	newTransaction := &Transaction{
		localStore: make(map[string]string),
	}

	newTransaction.next = ts.topTransaction
	ts.topTransaction = newTransaction
	ts.size++
}

func (ts *TransactionStack) Pop() {
	if ts.topTransaction == nil {
		fmt.Printf(MsgError + "No Active Transactions\n")
		return
	}

	ts.topTransaction = ts.topTransaction.next
	ts.size--
}

func (ts *TransactionStack) Rollback() {
	if ts.topTransaction == nil {
		fmt.Printf(MsgError + "No Active Transactions\n")
		return
	}

	for k := range ts.topTransaction.localStore {
		delete(ts.topTransaction.localStore, k)
	}
}

func (ts *TransactionStack) Set(key string, value string) {
	activeTransaction := ts.GetTopTransaction()

	if activeTransaction != nil {
		activeTransaction.localStore[key] = value
		return
	}

	if err := db.DB.Put([]byte(key), []byte(value), nil); err != nil {
		fmt.Printf(MsgError + err.Error() + "\n")
	}
}
