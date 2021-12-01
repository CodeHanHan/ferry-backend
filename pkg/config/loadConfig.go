package config

import "github.com/spf13/viper"

const (
	ConfigPath = "deploy/config/"
	ConfigName = "setting"
	ConfigType = "yml"
)

type Config struct {
	Version     string    `mapstructure:"version"`
	Mode        string    `mapstructure:"mode"`
	Application AppConfig `mapstructure:"application"`
	Database    DBConfig  `mapstructure:"database"`
	Jwt         JWTConfig `mapstructure:"jwt"`
	Author      []string  `mapstructure:"author"`
}

type AppConfig struct {
	ServerAddress string `mapstructure:"server_address"`
}

type DBConfig struct {
	DBDriver    string `mapstructure:"driver"`
	LoggerLevel int    `mapstructure:"level"`
	DBUser      string `mapstructure:"user"`
	DBPassword  string `mapstructure:"password"`
	DBPort      int    `mapstructure:"port"`
	DBName      string `mapstructure:"name"`
	DBHost      string `mapstructure:"host"`
	ParseTime   bool   `mapstructure:"parse_time"`
}

type JWTConfig struct {
	Secret  string `mapstructure:"secret"`
	Timeout string `mapstructure:"timeout"`
}

func SetUp(register func(*Config)) error {
	viper.AddConfigPath(ConfigPath)
	viper.SetConfigName("setting")
	viper.SetConfigType("yml")

	var config *Config = &Config{}
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	register(config)
	return nil
}
