package config

type OptionFunc func(*Config)

// WithExtension option function for stating config file extension
func WithExtension(ext string) OptionFunc {
	return func(cfg *Config) {
		cfg.extension = ext
	}
}

// WithFilename option function for stating config file name
func WithFilename(name string) OptionFunc {
	return func(cfg *Config) {
		cfg.name = name
	}
}

// WithPath option function for stating config file path
func WithPath(path string) OptionFunc {
	return func(cfg *Config) {
		cfg.path = path
	}
}
