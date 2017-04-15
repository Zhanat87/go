package app

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/go-ozzo/ozzo-validation"
)

// Config stores the application-wide configurations
var Config appConfig

type appConfig struct {
	// the path to the error message file. Defaults to "config/errors.yaml"
	ErrorFile             string `mapstructure:"error_file"`
	// the server port. Defaults to 8080
	ServerPort            int    `mapstructure:"server_port"`
	// the data source name (DSN) for connecting to the database. required.
	DSN                   string `mapstructure:"dsn"`
	DSN_DOCKER            string `mapstructure:"dsn_docker"`
	DSN_DOCKER_COMPOSE_V3 string `mapstructure:"dsn_docker_compose_v3"`
	// the signing method for JWT. Defaults to "HS256"
	JWTSigningMethod      string `mapstructure:"jwt_signing_method"`
	// JWT signing key. required.
	JWTSigningKey         string `mapstructure:"jwt_signing_key"`
	// JWT verification key. required.
	JWTVerificationKey    string `mapstructure:"jwt_verification_key"`
}

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.DSN, validation.Required),
		validation.Field(&config.DSN_DOCKER, validation.Required),
		validation.Field(&config.JWTSigningKey, validation.Required),
		validation.Field(&config.JWTVerificationKey, validation.Required),
	)
}

func (config appConfig) GetDSN() string {
	//_, issetDocker := os.LookupEnv("POSTGRESQL_PORT")
	//if issetDocker {
	//	return fmt.Sprintf(config.DSN_DOCKER, os.Getenv("POSTGRESQL_ENV_POSTGRES_USER"),
	//		os.Getenv("POSTGRESQL_ENV_POSTGRES_PASSWORD"), os.Getenv("POSTGRESQL_PORT_5432_TCP_ADDR"),
	//		os.Getenv("POSTGRESQL_PORT_5432_TCP_PORT"), os.Getenv("POSTGRESQL_ENV_POSTGRES_DB"))
	//}
	return config.DSN
	//return config.DSN_DOCKER_COMPOSE_V3
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
// Environment variables with the prefix "RESTFUL_" in their names are also read automatically.
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("restful")
	v.AutomaticEnv()
	v.SetDefault("error_file", "config/errors.yaml")
	v.SetDefault("server_port", 8080)
	v.SetDefault("jwt_signing_method", "HS256")
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return Config.Validate()
}
