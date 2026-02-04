package config

type Config struct {
	Port        int    `env:"PORT" envDefault:"3000"`
	AllowOrigin string `env:"CORS_ALLOW_ORIGIN" envDefault:"*"`
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://user:password@localhost/db"`
}
