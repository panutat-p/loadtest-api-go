package main

import (
	"fmt"
	"sync/atomic"

	"github.com/caarlos0/env"
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"

	"loadtest-api-go/internal"
)

func main() {
	var conf internal.Config
	err := env.Parse(&conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("conf: %+v\n", conf)

	var counter atomic.Uint64
	r1 := rate.NewLimiter(1, 5)

	h := internal.NewHandler(&counter, r1)

	e := echo.New()
	e.GET("/fruits", h.ListFruits)

	conf.Port = 8080
	err = e.Start(fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		panic(err)
	}
}
