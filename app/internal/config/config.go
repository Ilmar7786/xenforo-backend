package config

import (
	"context"
	"flag"
	"os"
	"sync"
	"time"

	"xenforo/app/pkg/logging"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		IsDebug bool `yaml:"is-debug" env-default:"false"`
		Jwt     struct {
			AccessTokenPrivateKey  string        `yaml:"ACCESS_TOKEN_PRIVATE_KEY"`
			AccessTokenExpiredIn   time.Duration `yaml:"ACCESS_TOKEN_EXPIRED_IN"`
			AccessTokenMaxAge      int           `yaml:"ACCESS_TOKEN_MAX_AGE"`
			RefreshTokenPrivateKey string        `yaml:"REFRESH_TOKEN_PRIVATE_KEY"`
			RefreshTokenExpiredIn  time.Duration `yaml:"REFRESH_TOKEN_EXPIRED_IN"`
			RefreshTokenMaxAge     int           `yaml:"REFRESH_TOKEN_MAX_AGE"`
		} `yaml:"jwt-token"`
	} `yaml:"app"`
	HTTP struct {
		IP           string        `yaml:"ip" env:"HTTP_IP" env-default:"0.0.0.0"`
		Port         uint          `yaml:"port" env:"HTTP_PORT" env-default:"8080"`
		ReadTimeout  time.Duration `yaml:"read-timeout" env:"HTTP_READ_TIMEOUT"`
		WriteTimeout time.Duration `yaml:"write-timeout" env:"HTTP_WRITE_TIMEOUT"`
		CORS         struct {
			AllowedMethods     []string `yaml:"allowed_methods" env:"HTTP-CORS-ALLOWED-METHODS"`
			AllowedOrigins     []string `yaml:"allowed_origins"`
			AllowCredentials   bool     `yaml:"allow_credentials"`
			AllowedHeaders     []string `yaml:"allowed_headers"`
			OptionsPassthrough bool     `yaml:"options_passthrough"`
			ExposedHeaders     []string `yaml:"exposed_headers"`
			Debug              bool     `yaml:"debug"`
		} `yaml:"cors"`
	} `yaml:"http"`
	PostgreSQL struct {
		Username string `yaml:"username" env:"PSQL_USER" env-required:"true"`
		Password string `yaml:"password" env:"PSQL_PASSWORD" env-required:"true"`
		Host     string `yaml:"host" env:"PSQL_HOST" env-required:"true"`
		Port     string `yaml:"port" env:"PSQL_PORT" env-required:"true"`
		Database string `yaml:"database" env:"PSQL_DATABASE" env-required:"true"`
	} `yaml:"postgresql"`
	Mail struct {
		From     string `yaml:"from" env:"MAIL_FROM" emv-required:"true"`
		Password string `yaml:"password" env:"MAIL_PASSWORD" emv-required:"true"`
		Username string `yaml:"username" env:"MAIL_USERNAME" emv-required:"true"`
		Host     string `yaml:"host" env:"MAIL_HOST" emv-required:"true"`
		Port     int    `yaml:"port" env:"MAIL_PORT" emv-required:"true"`
	} `yaml:"mail"`
}

const (
	EnvConfigPathName  = "CONFIG-PATH"
	FlagConfigPathName = "config"
)

var configPath string
var instance *Config
var once sync.Once

func GetConfig(ctx context.Context) *Config {
	once.Do(func() {
		flag.StringVar(&configPath, FlagConfigPathName, "configs/config.local.yaml", "this is app config file")
		flag.Parse()

		if err := godotenv.Load(); err != nil {
			logging.Warning(ctx, "virtual environment file (.env) not found")
		}

		logging.Info(ctx, "config init")

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			logging.Fatal(ctx, "config path is required")
		}

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logging.Info(ctx, help)
			logging.Fatal(ctx, err)
		}
	})
	return instance
}
