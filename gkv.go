package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gusandrioli/gkvDB/db"
	"github.com/pkg/errors"
)

var (
	gkv = &gkvDB{}
)

// Wraps the entire gkvDB app
type gkvDB struct {
	mainDB         *db.LevelDB
	helperDB       *db.LevelDB
	topTransaction *Transaction
	size           int
}

type Transaction struct {
	localStore map[string]string
	next       *Transaction
}

func (gkv *gkvDB) Commit() {
	activeTransaction := gkv.GetTopTransaction()

	if activeTransaction == nil {
		fmt.Printf(MsgWarning + "Nothing to commit")
		return
	}

	for k, v := range activeTransaction.localStore {
		if err := gkv.mainDB.DB.Put([]byte(k), []byte(v), nil); err != nil {
			fmt.Printf(MsgError + err.Error() + "\n")
		}

		if activeTransaction.next != nil {
			activeTransaction.next.localStore[k] = v
		}
	}
}

func (gkv *gkvDB) CountRecords() {
	var i int64
	iter := gkv.mainDB.DB.NewIterator(nil, nil)
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

func (gkv *gkvDB) DeleteRecord(key string) {
	activeTransaction := gkv.GetTopTransaction()

	if activeTransaction != nil {
		if _, ok := activeTransaction.localStore[key]; ok {
			delete(activeTransaction.localStore, key)
			return
		}
	}

	if err := gkv.mainDB.DB.Delete([]byte(key), nil); err != nil {
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

func (gkv *gkvDB) Get(key string) {
	activeTransaction := gkv.GetTopTransaction()

	if activeTransaction != nil {
		if v, ok := activeTransaction.localStore[key]; ok {
			fmt.Printf("%s\n", v)
			return
		}
	}

	value, err := gkv.mainDB.DB.Get([]byte(key), nil)
	if err != nil {
		fmt.Printf(MsgError+"%s not found\n", key)
		return
	}

	fmt.Printf("%s\n", string(value))
}

func (gkv *gkvDB) GetTopTransaction() *Transaction {
	return gkv.topTransaction
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

func (gkv *gkvDB) ListRecords() {
	iter := gkv.mainDB.DB.NewIterator(nil, nil)
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

func (gkv *gkvDB) ListTransaction() {
	if gkv.topTransaction == nil {
		fmt.Printf("No changes in transaction\n")
		return
	}

	transactionStackChanges := []map[string]string{}

	currentTransaction := gkv.topTransaction

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

func NewGkv(mainDB *db.LevelDB, helperDB *db.LevelDB) *gkvDB {
	return &gkvDB{
		mainDB:   mainDB,
		helperDB: helperDB,
	}
}

func (gkv *gkvDB) Push() {
	newTransaction := &Transaction{
		localStore: make(map[string]string),
	}

	newTransaction.next = gkv.topTransaction
	gkv.topTransaction = newTransaction
	gkv.size++
}

func (gkv *gkvDB) Pop() {
	if gkv.topTransaction == nil {
		fmt.Printf(MsgError + "No Active Transactions\n")
		return
	}

	gkv.topTransaction = gkv.topTransaction.next
	gkv.size--
}

func (gkv *gkvDB) Rollback() {
	if gkv.topTransaction == nil {
		fmt.Printf(MsgError + "No Active Transactions\n")
		return
	}

	for k := range gkv.topTransaction.localStore {
		delete(gkv.topTransaction.localStore, k)
	}
}

func (gkv *gkvDB) Set(key string, value string) {
	activeTransaction := gkv.GetTopTransaction()

	if activeTransaction != nil {
		activeTransaction.localStore[key] = value
		return
	}

	if err := gkv.mainDB.DB.Put([]byte(key), []byte(value), nil); err != nil {
		fmt.Printf(MsgError + err.Error() + "\n")
	}
}

func (gkv *gkvDB) SetExpire(key string, value string) {
	if err := gkv.helperDB.DB.Put(
		[]byte(gkv.mainDB.Name+"."+key),
		[]byte(value),
		nil,
	); err != nil {
		fmt.Printf(MsgError + err.Error() + "\n")
	}
}
