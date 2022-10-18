package storage

import (
	"fmt"
	"log"

	"github.com/alisaleemh/kv-service/config"
	"github.com/boltdb/bolt"
)

type IStorage interface {
	Delete(key string) error
	Insert(key string, value []byte) error
	Get(key string) []byte
}

type Storage struct {
	Db IStorage
}

// Currently new storage only returns a dynamo implementation

func NewStorage(envVars *config.EnvironmentVariables) *Storage {

	return &Storage{
		Db: &Bolt{
			db:         NewBoltStorage(),
			bucketName: "a",
		},
	}
}

func NewBoltStorage() *bolt.DB {

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
		panic("Cannot open db")
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("a"))
		if err != nil {
			return fmt.Errorf("couldn't create bucket: %s", err)
		}
		return nil
	})

	return db
}
