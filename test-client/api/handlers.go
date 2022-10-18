package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrReturnBody struct {
	Message string `json:"message"`
}

type GetReturnBody struct {
	Value string `json:"message"`
}

func (s *Server) HandleTestDelete() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		err := s.Service.TestDelete()

		if err != nil {
			err := fmt.Errorf("Error testing delete: %v", err)
			writeErr(w, r, err, http.StatusInternalServerError)
			return
		}

		resp := ErrReturnBody{
			Message: "Test Successful",
		}
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			err := fmt.Errorf("error unmarshalling: %v", err)
			writeErr(w, r, err, http.StatusInternalServerError)
			return
		}
		write(w, jsonResp)
	}
}

func (s *Server) HandleTestOverwrite() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		err := s.Service.TestOverwrite()

		if err != nil {
			err := fmt.Errorf("Error testing overwrite: %v", err)
			writeErr(w, r, err, http.StatusInternalServerError)
			return
		}

		resp := ErrReturnBody{
			Message: "Test Successful",
		}
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			err := fmt.Errorf("error unmarshalling: %v", err)
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
