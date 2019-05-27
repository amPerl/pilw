package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func request(key, path, method string) (string, error) {
	fullPath := fmt.Sprintf("%s%s", "https://app.pilw.io:8443/v1/", path)

	client := &http.Client{}

	req, err := http.NewRequest(method, fullPath, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("apikey", key)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func get(key, path string) (string, error) {
	result, err := request(key, path, "GET")
	return result, err
}
