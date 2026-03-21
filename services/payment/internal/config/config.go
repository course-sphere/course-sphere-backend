package config

type Config struct {
	Port        int    `env:"PORT" envDefault:"3000"`
	CorsOrigin  string `env:"CORS_ORIGIN" envDefault:"*"`
	ApiURL      string `env:"API_URL" envDefault:"http://localhost:8080"`
	ProxyURL    string `env:"PROXY_URL" envDefault:"http://localhost:8080"`
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://user:password@localhost/db"`

	PayOSClientID    string `env:"PAYOS_CLIENT_ID"`
	PayOSApiKey      string `env:"PAYOS_API_KEY"`
	PayOSChecksumKey string `env:"PAYOS_CHECKSUM_KEY"`
}
