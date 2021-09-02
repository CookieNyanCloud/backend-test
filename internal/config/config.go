package config

import (
	"github.com/spf13/viper"
	"time"
)

//значкения по умолчанию

var (
	defaults = map[string]interface{}{
		"defaultHttpPort"               : "8000",
		"defaultHttpRWTimeout"          : 10 * time.Second,
		"defaultHttpMaxHeaderMegabytes" : 1,
		"defaultLimiterRPS"             : 10,
		"defaultLimiterBurst"           : 2,
		"defaultLimiterTTL"             : 10 * time.Minute,
	}
)

//конфигурации
type (
	Config struct {
		Postgres PostgresConfig
		HTTP        HTTPConfig
		Limiter     LimiterConfig
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

func Init(configDir string) (*Config, error)  {

	//заполнение значений по умолчанию
	for k, v := range defaults{
		viper.SetDefault(k,v)
	}

	if err := parseConfigFile(configDir, viper.GetString("env")); err != nil {
		return nil, err
	}
	var cfg Config
	//if err := unmarshal(&cfg); err != nil {
	//	return nil, err
	//}

	if err := parseEnvVariables(); err != nil {
		return nil, err
	}
	return &cfg, nil
}



func parseEnvVariables() error {
	viper.SetEnvPrefix("postgres")
	return viper.BindEnv("pass")
}

func parseConfigFile(folder, env string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.SetConfigName(env)
	return viper.MergeInConfig()
}




























