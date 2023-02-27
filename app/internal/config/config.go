package config

import (
	"context"
	"flag"
	"log"
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
			AccessTokenPrivateKey  string        `yaml:"access-token-key"`
			AccessTokenExpiredIn   time.Duration `yaml:"access-token-expired-in"`
			RefreshTokenPrivateKey string        `yaml:"refresh-token-key"`
			RefreshTokenExpiredIn  time.Duration `yaml:"refresh-token-expired-in"`
		} `yaml:"jwt-token"`
	} `yaml:"app"`
	HTTP struct {
		IP           string        `yaml:"ip" env:"HTTP_IP"`
		Port         uint          `yaml:"port" env:"HTTP_PORT"`
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
	Mail struct {
		From     string `env:"MAIL_FROM" env-required:"true"`
		Password string `env:"MAIL_PASSWORD" env-required:"true"`
		Username string `env:"MAIL_USERNAME" env-required:"true"`
		Host     string `env:"MAIL_HOST" env-required:"true"`
		Port     int    `env:"MAIL_PORT" env-required:"true"`
		SSL      bool   `yaml:"ssl" env:"MAIL_SSL"`
	} `yaml:"mail"`
	ClientURL  string `yaml:"client-url" env:"CLIENT_URL" env-required:"true"`
	PostgreSQL struct {
		Username string `env:"PSQL_USER" env-required:"true"`
		Password string `env:"PSQL_PASSWORD" env-required:"true"`
		Host     string `env:"PSQL_HOST" env-required:"true"`
		Port     string `env:"PSQL_PORT" env-required:"true"`
		Database string `env:"PSQL_DATABASE" env-required:"true"`
	}
	FlashLiveSports struct {
		Token string `env:"FLASH_LIVE_SPORTS_TOKEN" env-required:"true"`
	}
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
			log.Println(help)
			log.Fatal(err)
		}
	})
	return instance
}
