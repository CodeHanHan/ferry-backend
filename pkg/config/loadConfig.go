package config

import "github.com/spf13/viper"

const (
	ConfigPath = "deploy/config/"
	ConfigName = "setting"
	ConfigType = "yml"
)

type Config struct {
	Version     string    `mapstructure:"version"`
	Application AppConfig `mapstructure:"application"`
	Database    DBConfig  `mapstructure:"database"`
}

type AppConfig struct {
	ServerAddress string `mapstructure:"server_address"`
}

type DBConfig struct {
	DBDriver   string `mapstructure:"db_driver"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBPort     int    `mapstructure:"db_port"`
	DBName     string `mapstructure:"db_name"`
	DBHost     string `mapstructure:"db_host"`
}

func LoadConfig() (config *Config, err error) {
	viper.AddConfigPath(ConfigPath)
	viper.SetConfigName("setting")
	viper.SetConfigType("yml")

	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err = viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	return
}
