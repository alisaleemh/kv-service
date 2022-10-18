package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ErrReturnBody struct {
	Message string `json:"message"`
}

type GetReturnBody struct {
	Value string `json:"message"`
}

func (s *Server) HandleGetKey() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		key, ok := vars["key"]
		if !ok {
			err := fmt.Errorf("missing key in request")
			writeErr(w, r, err, http.StatusBadRequest)
			return

		}

		val, err := s.Service.GetKey(key)

		if err != nil {
			err := fmt.Errorf("cannot get key: %v", err)
			writeErr(w, r, err, http.StatusInternalServerError)
			return
		}

		resp := GetReturnBody{
			Value: string(val),
		}

		jsonResp, err := json.Marshal(resp)

		if err != nil {
			err := fmt.Errorf("error unmarshalling val: %v", err)
			writeErr(w, r, err, http.StatusInternalServerError)
			return
		}

		write(w, jsonResp)
	}
}

func (s *Server) HandleDeleteKey() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		key, ok := vars["key"]
		if !ok {
			err := fmt.Errorf("missing key in request")
			writeErr(w, r, err, http.StatusBadRequest)
			return

		}

		err := s.Service.DeleteKey(key)

		if err != nil {
			err := fmt.Errorf("cannot delete key: %v", err)
			writeErr(w, r, err, http.StatusInternalServerError)
			return
		}

		resp := ErrReturnBody{
			Message: fmt.Sprintf("Deleted key %v", key),
		}
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			err := fmt.Errorf("error unmarshalling DeleteKeyBody: %v", err)
			writeErr(w, r, err, http.StatusInternalServerError)
			return
		}
		write(w, jsonResp)
	}
}

func (s *Server) HandleInsertKey() http.HandlerFunc {

	type InsertKeyBody struct {
		Value string `json:"value"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		key, ok := vars["key"]
		if !ok {
			err := fmt.Errorf("missing key in request")
			writeErr(w, r, err, http.StatusBadRequest)
			return

		}
		decoder := json.NewDecoder(r.Body)
		var insertKeyBody InsertKeyBody
		err := decoder.Decode(&insertKeyBody)
		if err != nil {
			err := fmt.Errorf("cannot decode %v, into insertKeyBody: %v", r.Body, err)
			writeErr(w, r, err, http.StatusInternalServerError)
			return
		}

		err = s.Service.InsertKey(key, []byte(insertKeyBody.Value))
		if err != nil {
			err := fmt.Errorf("insert Key failed: %v", err)
			writeErr(w, r, err, http.StatusInternalServerError)
			return
		}

		resp := ErrReturnBody{
			Message: fmt.Sprintf("Inserted key %v", key),
		}
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			err := fmt.Errorf("error unmarshalling return body: %v", err)
			writeErr(w, r, err, http.StatusInternalServerError)
			return
		}
		write(w, jsonResp)

	}
}

func write(w http.ResponseWriter, jsonResp []byte) {

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)

}

func writeErr(w http.ResponseWriter, r *http.Request, err error, code int) error {
	errObj := struct {
		Error string `json:"error"`
	}{Error: err.Error()}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	err = json.NewEncoder(w).Encode(errObj)
	if err != nil {
		return err
	}

	return nil
}
