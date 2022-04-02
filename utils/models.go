package utils

import "time"

type response struct {
	Message       string      `json:"message,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	AccessTimeIn  time.Time   `json:"access_time_in,omitempty"`
	AccessTimeOut time.Time   `json:"access_time_out,omitempty"`
}
