package internal

type Config struct {
	Port  int `env:"PORT"`
	Rate  int `env:"RATE"`
	Burst int `env:"BURST"`
}
