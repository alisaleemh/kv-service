package storage

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type Bolt struct {
	db         *bolt.DB
	bucketName string
}

func (b *Bolt) Insert(key string, value []byte) error {

	return b.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(b.bucketName))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), value)
		return err
	})
}

func (b *Bolt) Delete(key string) error {

	return b.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(b.bucketName))
		err := b.Delete([]byte(key))
		return err
	})
}

func (b *Bolt) Get(key string) []byte {
	var val []byte

	err := b.db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(b.bucketName))
		if bkt == nil {
			return fmt.Errorf("bucket %q not found", b.bucketName)
		}
		val = bkt.Get([]byte(key))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return val

}
