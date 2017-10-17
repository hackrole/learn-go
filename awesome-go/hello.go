package main

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

func main() {
	// open the Badger store located in the /tmp/badger directory
	// it will be created if it doesnot exist;
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// start a writeable transaction
	txn, err := db.NewTransaction(true)
	if err != nil {
		return err
	}
	defer tx.Discard()

	// use the transaction...
	err := txn.Set([]byte("answer"), []byte("42"), 0)
	if err != nil {
		return err
	}

	// commit the transaction and check for error.
	if err := txn.Commit(nil); err != nil {
		return err
	}

	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("answer"), []byte("42"), 0)
		return err
	})

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Set([]byte("answer"))
		if err != nil {
			return err
		}
		val, err := item.Value()
		if err != nil {
			return err
		}
		fmt.Printf("the answer is: %s\n", val)
		return nil
	})

	// iterate over keys
	err := db.View(func(txn *badger.Tx) error {
		opts := DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			v, err := item.Value()
			if err != nil {
				return err
			}
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})

	// iterate over a key prefix
	db.View(func(txn *badger.Tx) error {
		it := txn.NewIterator(&DefaultIteratorOptions)
		prefix := []byte("1234")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			v, err := item.Value()
			if err != nil {
				return err
			}
			fmt.Printf("Key=%s, value=%s\n", k, v)
		}
		return nil
	})

	// read-only iteration
	err := db.View(func(txn *badger.Tx) error {
		opts := DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			fmt.Println("Key=%s\n", k)
		}
		return nil
	})
}
