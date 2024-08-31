package main

import (
	"fmt"
	"sync/atomic"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"loadtest-api-go/internal"
	"loadtest-api-go/pkg"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var conf internal.Config
	err = env.Parse(&conf)
	if err != nil {
		panic(err)
	}
	pkg.PrintJSON(conf)
	conf.Validate()

	var counter atomic.Uint64
	h := internal.NewHandler(&counter)
	limiter := internal.RateLimiter(&counter, conf.Rate, conf.Burst)

	e := echo.New()
	e.GET("/", h.Health)
	e.GET("/fruits", h.ListFruits, limiter)

	err = e.Start(fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		panic(err)
	}
}
