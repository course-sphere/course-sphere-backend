package config

type Config struct {
	Port        int    `env:"PORT" envDefault:"3000"`
	CorsOrigin  string `env:"CORS_ORIGIN" envDefault:"*"`
	ProxyURL    string `env:"PROXY_URL" envDefault:"http://localhost:8080"`
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://user:password@localhost/db"`
}
