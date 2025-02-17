package config

import (
	"github.com/spf13/viper"
	"log/slog"
)

type AppConfig struct {
	Env  string `mapstructure:"APP_ENV"`
	Name string `mapstructure:"APP_NAME"`
}

type DatabaseConfig struct {
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
}

type AuthConfig struct {
	JwtSecret string `mapstructure:"JWT_SECRET"`
}

type RabbitConfig struct {
	RabbitHost string `mapstructure:"RABBIT_HOST"`
}

type RedisConfig struct {
	RedisHost string `mapstructure:"REDIS_HOST"`
}

type ServerConfig struct {
	APIDocsPort int    `mapstructure:"HTTP_PORT_API_DOCS"`
	Port        string `mapstructure:"HTTP_PORT"`
}

type Config struct {
	App      AppConfig      `mapstructure:",squash"`
	Database DatabaseConfig `mapstructure:",squash"`
	Auth     AuthConfig     `mapstructure:",squash"`
	Rabbit   RabbitConfig   `mapstructure:",squash"`
	Redis    RedisConfig    `mapstructure:",squash"`
	Server   ServerConfig   `mapstructure:",squash"`
}

var LoadedConfig Config

func LoadConfig(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("[APP]", "error", "cannot initialize configuration")
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		slog.Error("[APP]", "error", "cannot initialize configuration")
		panic(err)
	}

	slog.Info("[APP]", "message", "config was loaded successfully")

	LoadedConfig = config

	return config
}
