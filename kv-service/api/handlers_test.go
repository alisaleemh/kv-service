package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alisaleemh/kv-service/config"
	"github.com/alisaleemh/kv-service/service"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/alisaleemh/kv-service/storage"
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
func TestHandleGet(t *testing.T) {

	envVars := config.GetEnvVars()
	svc := service.Service{
		Storage: getTestDatabase(),
	}
	srv := Server{
		Service: svc,
		EnvVars: envVars,
	}
	req := httptest.NewRequest(http.MethodGet, "/ali", nil)
	req = mux.SetURLVars(req, map[string]string{
		"key": "test",
	})
	w := httptest.NewRecorder()

	handler := srv.HandleGetKey()
	handler.ServeHTTP(w, req)
	fmt.Print(w.Result().Body)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK, "Get Key handler status is %v", w.Result().StatusCode)

}

func TestHandleDelete(t *testing.T) {

	envVars := config.GetEnvVars()
	svc := service.Service{
		Storage: getTestDatabase(),
	}
	srv := Server{
		Service: svc,
		EnvVars: envVars,
	}
	req := httptest.NewRequest(http.MethodDelete, "/kv", nil)
	req = mux.SetURLVars(req, map[string]string{
		"key": "test",
	})
	w := httptest.NewRecorder()

	handler := srv.HandleDeleteKey()
	handler.ServeHTTP(w, req)
	fmt.Print(w.Result().Body)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK, "Get Key handler status is %v", w.Result().StatusCode)

}

func TestHandleInsert(t *testing.T) {

	envVars := config.GetEnvVars()
	svc := service.Service{
		Storage: getTestDatabase(),
	}
	srv := Server{
		Service: svc,
		EnvVars: envVars,
	}
	req := httptest.NewRequest(http.MethodPut, "/kv", strings.NewReader("{ \"value\" : \"tatti\" }"))
	req = mux.SetURLVars(req, map[string]string{
		"key": "test",
	})
	w := httptest.NewRecorder()

	handler := srv.HandleInsertKey()
	handler.ServeHTTP(w, req)
	fmt.Print(w.Result().Body)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK, "Get Key handler status is %v", w.Result().StatusCode)

}
