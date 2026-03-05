package config

type Config struct {
	Port       int    `env:"PORT" envDefault:"3000"`
	S3Endpoint string `env:"S3_ENDPOINT"`
	S3Bucket   string `env:"S3_BUCKET" envDefault:"course-sphere"`
}
