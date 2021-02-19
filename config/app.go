package config

type AppConfig struct {
	Name     string `env:"NAME" envDefault:"Sullivan"`
	Env      string `env:"ENV" envDefault:"production"`
	Port     int    `env:"PORT" envDefault:"8090"`
	Debug    bool   `env:"DEBUG" envDefault:"true"`
	Url      string `env:"URL"  envDefault:"http://localhost"`
	TimeZone string `env:"TIMEZONE"  envDefault:"UTC"`
	Locale   string `env:"LOCALE"  envDefault:"en"`
	Key      string `env:"key"`
}
