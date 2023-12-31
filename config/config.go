package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port int `envconfig:"PORT" default:"8085" required:"true"`

	DBHost         string `envconfig:"DB_HOST" default:"localhost"`
	DBPort         string `envconfig:"DB_PORT" default:"3307"`
	DBUserName     string `envconfig:"DB_USERNAME" default:"dev"`
	DBPassword     string `envconfig:"DB_PASSWORD" default:"dev"`
	DBDatabaseName string `envconfig:"DB_DBNAME" default:"books_db"`
	DBLogMode      int    `envconfig:"DB_LOG_MODE" default:"3"`
}

func New() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg

}
