package config

import (
	"github.com/spf13/viper"
	"time"
)

//конфигурации
type (
	Config struct {
		Postgres PostgresConfig
		HTTP     HTTPConfig
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

	LimiterConfig struct {
		RPS   int
		Burst int
		TTL   time.Duration
	}
)

func Init(configDir string) (*Config, error) {

	//чтение данных из файла конфигураций
	viper.AddConfigPath(configDir)
	viper.SetConfigName("main")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	//заполняем структуру значениями из файла конфигураций
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	// чтение .env файла
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	//заполняем структуру значениями из .env
	cfg.Postgres.Password = viper.GetString("postgres_password")
	cfg.ApiKey = viper.GetString("api_key")
	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	return nil
}
