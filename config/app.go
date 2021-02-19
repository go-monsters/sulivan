package config

type AppConfig struct {
	Name string `env:"NAME"`
	Port int    `env:"PORT" envDefault:"8090"`
}
