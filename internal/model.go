package internal

import (
	"errors"
)

var ErrTooManyRequests = errors.New("too many requests")

type ResponseSuccess struct {
	Count uint64 `json:"count"`
	Data  any    `json:"data"`
}

type ResponseError struct {
	Count uint64  `json:"count"`
	Token float64 `json:"token"`
	Error error   `json:"error"`
}
