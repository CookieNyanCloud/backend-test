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

	if err := parseEnv(); err != nil {
		return nil, err
	}


}

func parseEnv() error {
	if err := parsePostgresEnvVariables(); err != nil {
		return err
	}
	return parseHostFromEnv()
}

func parsePostgresEnvVariables() error {
	viper.SetEnvPrefix("postgres")
	return viper.BindEnv("pass")
}
func parseHostFromEnv() error {
	viper.SetEnvPrefix("http")
	return viper.BindEnv("host")
}



























