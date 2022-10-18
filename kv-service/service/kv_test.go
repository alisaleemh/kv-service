package service

import (
	"testing"

	"github.com/alisaleemh/kv-service/config"
	"github.com/alisaleemh/kv-service/storage"
	"github.com/stretchr/testify/assert"
)

type TestDb struct {
}

func (t TestDb) Delete(key string) error {
	return nil
}

func (t TestDb) Get(key string) []byte {
	return []byte("test")
}

func (t TestDb) Insert(key string, value []byte) error {
	return nil
}

func getTestDatabase() *storage.Storage {
	return &storage.Storage{
		Db: TestDb{},
	}
}

func TestGetKey(t *testing.T) {

	s := Service{
		Storage: getTestDatabase(),
		EnvVars: config.GetEnvVars(),
	}

	key, err := s.GetKey("test")
	assert.Nilf(t, err, "error is not nil %v", err)

	assert.NotNil(t, key, "Key is nil")

}

func TestInsertKey(t *testing.T) {

	s := Service{
		Storage: getTestDatabase(),
		EnvVars: config.GetEnvVars(),
	}

	err := s.InsertKey("test", []byte("test"))
	assert.Nilf(t, err, "error is not nil %v", err)

}

func TestDeleteKey(t *testing.T) {

	s := Service{
		Storage: getTestDatabase(),
		EnvVars: config.GetEnvVars(),
	}

	err := s.DeleteKey("test")
	assert.Nilf(t, err, "error is not nil %v", err)

}
