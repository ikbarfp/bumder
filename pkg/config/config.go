package config

import (
	"github.com/spf13/viper"
)

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

// New constructor for creating new config. It returns interface Configurator,
// so it will contain all of its signature method
func New(opts ...OptionFunc) Configurator {
	cfg := defaultOptions

	for _, opt := range opts {
		opt(cfg)
	}

	return &ViperConfig{config: cfg, viperCfg: viper.New()}
}

// GetString return string with . (period) delimiter if your config had nested object like viper-style
func GetString(name string) string {
	return viperConf.getString(name)
}

// GetStrings return string slices with . (period) delimiter if your config had nested object like viper-style
func GetStrings(name string) []string {
	return viperConf.getStrings(name)
}

// GetInt return int with . (period) delimiter if your config had nested object like viper-style
func GetInt(name string) int {
	return viperConf.getInt(name)
}

// GetInts return int slices with . (period) delimiter if your config had nested object like viper-style
func GetInts(name string) []int {
	return viperConf.getInts(name)
}

// GetInt32 return int32 with . (period) delimiter if your config had nested object like viper-style
func GetInt32(name string) int32 {
	return viperConf.getInt32(name)
}

// GetFloat64 return float64 with . (period) delimiter if your config had nested object like viper-style
func GetFloat64(name string) float64 {
	return viperConf.getFloat64(name)
}

// GetBool return boolean with . (period) delimiter if your config had nested object like viper-style
func GetBool(name string) bool {
	return viperConf.getBool(name)
}
