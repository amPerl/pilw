package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func createRequest(key, path, method string, data url.Values) (*http.Request, error) {
	fullPath := fmt.Sprintf("%s%s", "https://app.pilw.io:8443/v1/", path)

	var bodyReader io.Reader
	if len(data) > 0 {
		bodyReader = strings.NewReader(data.Encode())
	}

	req, err := http.NewRequest(method, fullPath, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("apikey", key)

	return req, nil
}

func doRequest(client *http.Client, req *http.Request) (string, error) {
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bodyStr := string(body)

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Server returned status %d: %s", resp.StatusCode, bodyStr)
	}

	return bodyStr, nil
}

func plainRequest(key, path, method string) (string, error) {
	client := &http.Client{}

	req, err := createRequest(key, path, method, url.Values{})
	if err != nil {
		return "", err
	}

	result, err := doRequest(client, req)
	if err != nil {
		return "", err
	}

	return result, nil
}

func formRequest(key, path, method string, data url.Values) (string, error) {
	client := &http.Client{}

	req, err := createRequest(key, path, method, data)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	result, err := doRequest(client, req)
	if err != nil {
		return "", err
	}

	return result, nil
}

func get(key, path string) (string, error) {
	result, err := plainRequest(key, path, "GET")
	return result, err
}

func postForm(key, path string, data url.Values) (string, error) {
	result, err := formRequest(key, path, "POST", data)
	return result, err
}

func deleteForm(key, path string, data url.Values) (string, error) {
	result, err := formRequest(key, path, "DELETE", data)
	return result, err
}

func patchForm(key, path string, data url.Values) (string, error) {
	result, err := formRequest(key, path, "PATCH", data)
	return result, err
}
