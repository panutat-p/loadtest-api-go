package internal

import (
	"errors"
)

var ErrTooManyRequests = errors.New("too many requests")

type ResponseHealth struct {
	Count   uint64 `json:"count"`
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Code  int    `json:"code"`
	Count uint64 `json:"count"`
	Data  any    `json:"data"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
