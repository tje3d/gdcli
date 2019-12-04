package gdlib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// RequestOptions ...
type RequestOptions struct {
	method   string
	data     string
	language string
	version  string
	token    string
}

func sendRequest(path string, options RequestOptions) string {
	var req *http.Request

	client := &http.Client{}
	fullURL := generateURLFromPath(path)
	requestValues := url.Values{}
	appendRequestValues(options.data, &requestValues)

	// ─── CREATE REQUEST ─────────────────────────────────────────────────────────────

	if options.data == "" || options.method == "GET" {
		req, _ = http.NewRequest(options.method, fullURL, nil)
	} else {
		req, _ = http.NewRequest(options.method, fullURL, strings.NewReader(requestValues.Encode()))
	}

	// ─── BIND REQUEST VALUES ────────────────────────────────────────────────────────

	if options.method == "GET" {
		req.URL.RawQuery = requestValues.Encode()
	} else if options.method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(requestValues.Encode())))
	}

	// ─── ADD HEADERS ────────────────────────────────────────────────────────────────

	req.Header.Add("APPVERSION", "web")

	if options.language != "" {
		req.Header.Add("Accept-Language", options.language)
	} else {
		req.Header.Add("Accept-Language", "fa")
	}

	if options.version != "" {
		req.Header.Add("X-VERSION", options.version)
	}

	if options.token != "" {
		req.Header.Add("X-Token", options.token)
	}

	// ─── SEND REQUEST ───────────────────────────────────────────────────────────────

	response, err := client.Do(req)

	if err != nil {
		return err.Error()
	}

	body, _ := ioutil.ReadAll(response.Body)
	result := string(body)

	return result
}

func generateURLFromPath(path string) string {
	const url = "https://core.gap.im/v1/%s.json"
	return fmt.Sprintf(url, path)
}

func appendRequestValues(input string, requestValues *url.Values) {
	if input == "" {
		return
	}

	var data map[string]interface{}

	err := json.Unmarshal([]byte(input), &data)

	if err != nil {
		return
	}

	for key, val := range data {
		str, ok := val.(string)

		if ok {
			requestValues.Add(key, str)
		}
	}
}
