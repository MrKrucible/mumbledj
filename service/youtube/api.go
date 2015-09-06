/*
 * MumbleDJ
 * By Matthieu Grieger
 * service/youtube/api.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package youtube

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jmoiron/jsonq"
)

// PerformGetRequest does all the grunt work for a YouTube HTTPS GET request.
func PerformGetRequest(url string) (*jsonq.JsonQuery, error) {
	var jsonString string
	var response *http.Response
	var err error
	var body []byte

	if response, err = http.Get(url); err != nil {
		return nil, errors.New("An error occurred while receiving HTTPS GET response.")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		if response.StatusCode == 403 {
			return nil, errors.New("An invalid YouTube API key has been supplied.")
		}
		return nil, errors.New("YouTube API request failed with status code " + string(response.StatusCode) + ".")
	}
	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, errors.New("An error occurred while reading YouTube API response.")
	}

	jsonString = string(body)
	jsonData := map[string]interface{}{}
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	decoder.Decode(&jsonData)
	jq := jsonq.NewQuery(jsonData)

	return jq, nil
}
