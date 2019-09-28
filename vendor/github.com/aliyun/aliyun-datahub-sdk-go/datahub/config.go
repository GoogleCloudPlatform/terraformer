package datahub

type Config struct {
	UserAgent string
}

func NewDefaultConfig() *Config {
	return &Config{
		UserAgent: DefaultUserAgent(),
	}
}
