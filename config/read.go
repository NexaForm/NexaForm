package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

func absPath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	return filepath.Abs(path)

}

func ReadGeneric[T any](configPath string) (T, error) {
	var config T

	fullAbsPath, err := absPath(configPath)
	if err != nil {
		return config, err
	}

	viper.SetConfigFile(fullAbsPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

func ReadStandard(configPath string) (Config, error) {
	return ReadGeneric[Config](configPath)
}
