package gdlib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	url := generateURLFromPath(path)

	if options.data == "" || options.method == "GET" {
		req, _ = http.NewRequest(options.method, url, nil)
	} else {
		req, _ = http.NewRequest(options.method, url, strings.NewReader(options.data))
	}

	if options.method == "GET" && options.data != "" {
		var data map[string]interface{}

		err := json.Unmarshal([]byte(options.data), &data)

		if err == nil {
			query := req.URL.Query()

			for key, val := range data {
				str, ok := val.(string)

				if ok {
					query.Add(key, str)
				}
			}

			req.URL.RawQuery = query.Encode()
		}
	}

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

	response, err := client.Do(req)

	if err != nil {
		return err.Error()
	}

	body, _ := ioutil.ReadAll(response.Body)
	result := string(body)

	return result
}

func generateURLFromPath(path string) string {
	return fmt.Sprintf("https://core.gap.im/v1/%s.json", path)
}
