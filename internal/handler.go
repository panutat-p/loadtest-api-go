package internal

import (
	"net/http"
	"sync/atomic"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	conf    *Config
	counter *atomic.Uint64
}

func NewHandler(conf *Config, counter *atomic.Uint64) *Handler {
	return &Handler{
		conf:    conf,
		counter: counter,
	}
}

func (h *Handler) Health(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		ResponseHealth{
			Count:   h.counter.Load(),
			Message: "running",
		},
	)
}

func (h *Handler) ListFruits(c echo.Context) error {
	n := h.counter.Add(1)

	time.Sleep(h.conf.DelayFruit)

	fruits := []string{
		"apple",
		"banana",
		"cherry",
		"date",
		"elderberry",
		"fig",
		"grape",
		"honeydew",
		"imbe",
		"jackfruit",
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
		ResponseSuccess{
			Code:  200,
			Count: n,
			Data:  fruits,
		},
	)
}
