package config

type Config struct {
	Port       int    `env:"PORT" envDefault:"3000"`
	CorsOrigin string `env:"CORS_ORIGIN" envDefault:"*"`
	S3Endpoint string `env:"S3_ENDPOINT"`
	S3Bucket   string `env:"S3_BUCKET" envDefault:"course-sphere"`
}
