package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

func ReturnResponse(w http.ResponseWriter, accessTimeIn time.Time, data interface{}, errorMsg string, responseNumber int) {
	response := response{
		Message:       errorMsg,
		Data:          data,
		AccessTimeIn:  accessTimeIn,
		AccessTimeOut: time.Now(),
	}
	resp, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseNumber)
	w.Write(resp)
}
