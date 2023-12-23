package config

import "github.com/spf13/viper"

type Configurator interface {
	Load() error
}

type Config struct {
	Name, Extension, Path string

	name, extension, path string
}

var (
	defaultOptions = &Config{
		name:      "config",
		extension: "yml",
		path:      ".",
	}
)

// New . . .
func New(opts ...OptionFunc) Configurator {
	cfg := defaultOptions

	for _, opt := range opts {
		opt(cfg)
	}

	return &ViperConfig{config: cfg, viperCfg: viper.New()}
}

// GetString . . .
func GetString(name string) string {
	return viperConf.getString(name)
}

// GetStrings . . .
func GetStrings(name string) []string {
	return viperConf.getStrings(name)
}

// GetInt . . .
func GetInt(name string) int {
	return viperConf.getInt(name)
}

// GetInts . . .
func GetInts(name string) []int {
	return viperConf.getInts(name)
}

// GetInt32 . . .
func GetInt32(name string) int32 {
	return viperConf.getInt32(name)
}

// GetFloat64 . . .
func GetFloat64(name string) float64 {
	return viperConf.getFloat64(name)
}

// GetBool . . .
func GetBool(name string) bool {
	return viperConf.getBool(name)
}
