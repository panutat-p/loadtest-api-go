package internal

import (
	"net/http"
	"sync/atomic"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

type Handler struct {
	counter *atomic.Uint64
	limiter *rate.Limiter
}

func NewHandler(counter *atomic.Uint64, r1 *rate.Limiter) *Handler {
	return &Handler{
		counter: counter,
		limiter: r1,
	}
}

type ResponseFruit struct {
	Count uint64   `json:"count"`
	Fruit []string `json:"fruit"`
}

func (h *Handler) ListFruits(c echo.Context) error {
	n := h.counter.Add(1)

	fruits := []string{
		"apple",
		"banana",
		"cherry",
		"date",
		"elderberry",
		"fig",
		"grape",
		"honeydew",
		"kiwi",
		"lemon",
		"mango",
		"nectarine",
		"orange",
		"papaya",
		"quince",
		"raspberry",
		"strawberry",
		"tangerine",
		"ugli",
		"vanilla",
		"watermelon",
		"xinomavro",
		"yuzu",
		"zucchini",
	}
	return c.JSON(
		http.StatusOK,
		ResponseFruit{
			Count: n,
			Fruit: fruits,
		},
	)
}
