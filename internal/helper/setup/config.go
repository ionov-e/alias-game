package setup

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv" //nolint:nolintlint,goimports
	"sync"
)

type Config struct {
	IsDebug  bool `yaml:"debug" env-default:"true"` //nolint:tagliatelle
	Telegram struct {
		Token string `yaml:"token" env-required:"true"`
	} `yaml:"telegram"`
	ConcurrencyLimit int `yaml:"concurrency-limit" env-required:"true"`
	Redis            struct {
		Address  string `yaml:"address" env-required:"true"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
	Webhook struct {
		Enabled bool   `yaml:"enabled" env-default:"false"`
		URL     string `yaml:"url"`
		Port    string `yaml:"port"`
	} `yaml:"webhook"`
}

var instance *Config
var once sync.Once
var initError error // Переменная для хранения ошибки инициализации

func GetConfig() (*Config, error) {
	once.Do(func() {
		instance = &Config{}
		err := cleanenv.ReadConfig("config.yml", instance)
		if err != nil {
			help, err2 := cleanenv.GetDescription(instance, nil)
			initError = fmt.Errorf("config error: %w; help: %s; err2: %w", err, help, err2)
		}
	})

	if initError != nil {
		return nil, initError // Возвращаем ошибку, если она была
	}
	return instance, nil
}
