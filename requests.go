package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func get(url string, out interface{}) error {
	// make get request to url
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	//read the body and convert to string
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, out)
}
