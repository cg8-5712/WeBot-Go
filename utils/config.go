package utils

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LoginMode     string        `env:"LOGIN_MODE" default:"normal" required:"true"`
	HotReloadFile string        `env:"HOT_RELOAD_FILE" default:"storage.json"`
	ReconnectSec  time.Duration `env:"BOT_RECONNECT_INTERVAL" default:"10s"`
	LogLevel      string        `env:"LOG_LEVEL" default:"info"`
}

var Cfg *Config

func LoadConfig() {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("配置加载失败: %v", err)
	}
	Cfg = &cfg
}
