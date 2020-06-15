package config

import (
	"errors"

	"github.com/spf13/viper"
)

// Init sets up the config for the app (using Viper)
func Init(v *viper.Viper) (*viper.Viper, error) {
	if v == nil {
		v = viper.New()
		v.SetConfigName("config")
		v.AddConfigPath(".")
		v.SetConfigType("yaml")

		err := v.ReadInConfig()
		if err != nil {
			return nil, err
		}
	}

	if GetDSN(v) == "" {
		return nil, errors.New("missing required config DSN")
	}

	return v, nil
}

// GetDSN returns the DSN from config
func GetDSN(v *viper.Viper) string {
	return v.GetString("DSN")
}
