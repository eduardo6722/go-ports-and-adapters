package handlers

import "encoding/json"

func jsonError(msg string) []byte {
	messageErr := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	result, err := json.Marshal(messageErr)
	if err != nil {
		return []byte(err.Error())
	}
	return result
}
