package internal

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
)

type Config struct {
	Port           int           `env:"PORT"`
	Rate           int           `env:"RATE"`
	Burst          int           `env:"BURST"`
	DelayRateLimit time.Duration `env:"DELAY_RATE_LIMIT"`
	DelayFruit     time.Duration `env:"DELAY_FRUIT"`
}

func (c *Config) Validate() {
	err := validation.ValidateStruct(c,
		validation.Field(&c.Port, validation.Required),
		validation.Field(&c.Rate, validation.Required),
		validation.Field(&c.Burst, validation.Required),
	)
	if err != nil {
		panic(err)
	}
}
