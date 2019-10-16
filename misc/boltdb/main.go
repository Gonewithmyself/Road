package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

/*
	mongodb  	boltdb
	db       	db
	collection	bucket
*/
func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		bk, _ := tx.CreateBucketIfNotExists([]byte("test"))

		bk.Put([]byte("go"), []byte("1"))
		bk.Put([]byte("python"), []byte("2"))
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte("test"))
		fmt.Println(string(bk.Get([]byte("python"))))
		return nil
	})
}

// C
// R
// U
// D
