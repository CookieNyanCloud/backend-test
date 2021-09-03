package config

import (
	"github.com/spf13/viper"
	"time"
)

//todo: mapstructure

//значкения по умолчанию
var (
	defaults = map[string]interface{}{
		"defaultHttpPort":               "8000",
		"defaultHttpRWTimeout":          10 * time.Second,
		"defaultHttpMaxHeaderMegabytes": 1,
		"defaultLimiterRPS":             10,
		"defaultLimiterBurst":           2,
		"defaultLimiterTTL":             10 * time.Minute,
	}
)

//конфигурации
type (
	Config struct {
		Postgres PostgresConfig
		HTTP     HTTPConfig
		Limiter  LimiterConfig
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

	//заполнение значений по умолчанию
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}

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
	cfg.Postgres.Password = viper.GetString("postgres_pass")
	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("limiter", &cfg.Limiter); err != nil {
		return err
	}
	return nil
}
