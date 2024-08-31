package internal

import (
	"net/http"
	"sync/atomic"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	counter *atomic.Uint64
}

func NewHandler(counter *atomic.Uint64) *Handler {
	return &Handler{
		counter: counter,
	}
}

func (h *Handler) Health(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		map[string]any{
			"message": "running",
		},
	)
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
		ResponseSuccess{
			Count: n,
			Data:  fruits,
		},
	)
}
