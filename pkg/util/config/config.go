package config

import (
	"strings"

	"github.com/frontierdigital/ranger/pkg/cmd/app"

	"github.com/spf13/viper"
)

func LoadConfig() (config *app.Config, err error) {
	viper.SetEnvPrefix("RANGER")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err = viper.BindEnv("ado.pat")
	if err != nil {
		return nil, err
	}

	err = viper.BindEnv("git.useremail")
	if err != nil {
		return nil, err
	}

	err = viper.BindEnv("git.username")
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)

	return config, err
}
