package utils

import (
	"encoding/json"
	"net/http"
)

func GetJson(response *http.Response, target interface{}) interface{} {
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&target)

	return &target
}

func ToJSON(data []byte, target interface{}) error {
	return json.Unmarshal(data, &target)
}
