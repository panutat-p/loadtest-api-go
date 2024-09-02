package internal

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

func RateLimiter(conf *Config, counter *atomic.Uint64, r, b int) echo.MiddlewareFunc {
	limiter := rate.NewLimiter(rate.Limit(r), b)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !limiter.Allow() {
				fmt.Printf("🔴 %v %v -> %v\n", c.Request().Method, c.Request().URL.Path, http.StatusTooManyRequests)
				time.Sleep(conf.DelayRateLimit)
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
