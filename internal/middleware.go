package internal

import (
	"net/http"
	"sync/atomic"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

func RateLimiter(counter *atomic.Uint64, r, b int) echo.MiddlewareFunc {
	limiter := rate.NewLimiter(rate.Limit(r), b)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !limiter.Allow() {
				return c.JSON(
					http.StatusTooManyRequests,
					ResponseError{
						Count: counter.Load(),
						Token: limiter.Tokens(),
						Error: ErrTooManyRequests,
					},
				)
			}
			return next(c)
		}
	}
}
