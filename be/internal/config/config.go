package config

import (
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres" //driver for postgres
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

//Config struct
type Config struct {
	Address  string `envconfig:"SERVER_HOST"`
	Port     string `envconfig:"SERVER_PORT"`
	Env      string `envconfig:"ENV"`
	Database DbConfig
}

//DefaultConfig returns default config
//used if no config file is found
func DefaultConfig() *Config {
	return &Config{
		Address:  "0.0.0.0",
		Port:     "5000",
		Env:      "dev",
		Database: DefaultDbConfig(),
	}
}

//DbConfig represents data useful to db connection
type DbConfig struct {
	Host     string `envconfig:"PSQL_HOST"`
	Port     string `envconfig:"PSQL_PORT"`
	User     string `envconfig:"PSQL_USER"`
	Password string `envconfig:"PSQL_PSW"`
	Name     string `envconfig:"PSQL_DATABASE"`
}

//Dialect returns mocked postgres dialect
func (c DbConfig) Dialect() string {
	return "postgres"
}

//ConnectionInfo returns string to connect to postgres db
func (c DbConfig) ConnectionInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
}

//DefaultDbConfig returns default config for connecting to postgres
func DefaultDbConfig() DbConfig {
	return DbConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "vagrant",
		Password: "vagrant",
		Name:     "vagrant",
	}
}

//Load returns config based on .config file if exists, else use default config
func Load() *Config {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		log.Errorf("Error loading configuration. Loaded Default Configuration %v", err.Error())
		return DefaultConfig()
	}

	log.Info("Config loaded successfully")
	return &c
}

//SetLogConfig initialization of logrus
func SetLogConfig() {

	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		PadLevelText:           true,
		DisableLevelTruncation: false,
	})
	log.SetReportCaller(true)

	log.Info("Setup Logging")
}
