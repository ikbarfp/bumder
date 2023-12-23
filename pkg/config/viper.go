package config

import "github.com/spf13/viper"

var viperConf *ViperConfig

type ViperConfig struct {
	config   *Config
	viperCfg *viper.Viper
}

func (v *ViperConfig) Load() error {
	v.viperCfg.SetConfigName(v.config.name)
	v.viperCfg.SetConfigType(v.config.extension)
	v.viperCfg.AddConfigPath(v.config.path)

	// Read all config
	if err := v.viperCfg.ReadInConfig(); err != nil {
		return err
	}

	v.viperCfg.AutomaticEnv()
	v.viperCfg.WatchConfig()

	viperConf = v

	return nil
}

func (v *ViperConfig) getString(name string) string {
	return v.viperCfg.GetString(name)
}

func (v *ViperConfig) getStrings(name string) []string {
	return v.viperCfg.GetStringSlice(name)
}

func (v *ViperConfig) getInt(name string) int {
	return v.viperCfg.GetInt(name)
}

func (v *ViperConfig) getInts(name string) []int {
	return v.viperCfg.GetIntSlice(name)
}

func (v *ViperConfig) getInt32(name string) int32 {
	return v.viperCfg.GetInt32(name)
}

func (v *ViperConfig) getFloat64(name string) float64 {
	return v.viperCfg.GetFloat64(name)
}

func (v *ViperConfig) getBool(name string) bool {
	return v.viperCfg.GetBool(name)
}
