package config

import (
	"fmt"

	"github.com/LightAlykard/SupportTgBot/internal/log"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramLoggerBotToken string `mapstructure:"bot_token"`
	TelegramAdminUserID    int64  `mapstructure:"admin_user_id"`
	DefaultBD              string `mapstructure:"default_bd"`
	BotAddress             string `mapstructure:"bot_address"`
	BotPort                string `mapstructure:"bot_port"`
	CertPath               string `mapstructure:"cert_path"`
	KeyPath                string `mapstructure:"key_path"`
}

func ReadConfig(path, configFile, configType string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configFile)
	viper.SetConfigType(configType)

	viper.SetDefault("bot_token", "")
	viper.SetDefault("admin_user_id", 0)
	viper.SetDefault("default_bd", "")
	viper.SetDefault("bot_address", "")
	viper.SetDefault("bot_port", "")
	viper.SetDefault("cert_path", "")
	viper.SetDefault("key_path", "")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Warn().Msg("used default configs")
	} else {
		log.Info().Msgf("loaded config from %s", configType)
	}

	err = viper.Unmarshal(&config)

	return config, err
}
