package internal

var ErrTooManyRequests = "Too many requests"

type ResponseSuccess struct {
	Count uint64 `json:"count"`
	Data  any    `json:"data"`
}

type ResponseError struct {
	Count uint64  `json:"count"`
	Token float64 `json:"token"`
	Error string  `json:"error"`
}
