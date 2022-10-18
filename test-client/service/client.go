package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type InsertKeyBody struct {
	Value string `json:"value"`
}

func BuildDeleteUrl(baseUrl string) string {
	return baseUrl + "/kv/test"
}

func (S *Service) TestDelete() error {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, BuildDeleteUrl(S.EnvVars.KV_SERVICE_URL), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}

	return nil

}

func (S *Service) TestOverwrite() error {

	origInsertBody := InsertKeyBody{
		Value: "testOrig",
	}
	overWriteBody := InsertKeyBody{
		Value: "testOverwrite",
	}
	client := &http.Client{}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(origInsertBody)

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/kv/test", payloadBuf)
	if err != nil {
		log.Println(err)
		return err
	}

	// Fetch Request
	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	payloadBuf = new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(overWriteBody)

	req, err = http.NewRequest(http.MethodPut, "http://localhost:8080/kv/test", payloadBuf)
	if err != nil {
		log.Println(err)
		return err
	}

	// Fetch Request
	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	type GetReturnBody struct {
		Value string `json:"value"`
	}

	req, err = http.NewRequest(http.MethodGet, "http://localhost:8080/kv/test", nil)
	if err != nil {
		log.Println(err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	var getBody GetReturnBody

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&getBody)
	if err != nil {
		log.Println(err)
		return err
	}

	if !strings.Contains(getBody.Value, overWriteBody.Value) {
		return fmt.Errorf("Test failed, over writen is not correct")
	}

	return nil

}
