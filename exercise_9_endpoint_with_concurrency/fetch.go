package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const Url = "https://xkcd.com"

func Fetch(n int) (*jsonRecieved, error) {
	client := &http.Client{Timeout: time.Minute * 5}

	var link []string = []string{Url, fmt.Sprintf("%d", n), "info.0.json"}
	url := strings.Join(link, "/")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data jsonRecieved

	if resp.StatusCode != http.StatusOK {
		data = jsonRecieved{}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, fmt.Errorf("json err: %v", err)
		}
	}
	resp.Body.Close()
	return &data, nil
}
