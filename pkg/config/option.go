package config

type OptionFunc func(*Config)

// WithExtension . . .
func WithExtension(ext string) OptionFunc {
	return func(cfg *Config) {
		cfg.extension = ext
	}
}

// WithFilename . . .
func WithFilename(name string) OptionFunc {
	return func(cfg *Config) {
		cfg.name = name
	}
}

// WithPath . . .
func WithPath(path string) OptionFunc {
	return func(cfg *Config) {
		cfg.path = path
	}
}
