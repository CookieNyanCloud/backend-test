package config

import (
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Postgres PostgresConfig
		HTTP     HTTPConfig
		Redis    RedisConfig
		ApiKey   string
	}

	PostgresConfig struct {
		Host     string
		Port     string
		Username string
		DBName   string
		SSLMode  string
		Password string
	}

	HTTPConfig struct {
		Host               string
		Port               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		MaxHeaderMegabytes int
	}
	RedisConfig struct {
		Addr string
	}

	LimiterConfig struct {
		RPS   int
		Burst int
		TTL   time.Duration
	}
)

func Init(configDir string, local bool) (*Config, error) {

	//reading yaml config file
	viper.AddConfigPath(configDir)
	viper.SetConfigName("main")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	//creating config struct
	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	//reading .env config file
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	cfg.Postgres.Password = viper.GetString("postgres_password")
	cfg.ApiKey = viper.GetString("api_key")
	if !local {
		cfg.Postgres.Host = viper.GetString("host")
	}
	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("redis", &cfg.Redis); err != nil {
		return err
	}
	return nil
}
