package config

import (
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

type Configlist struct {
	ApiKey      string
	ApiSecret   string
	LogFile     string
	ProductCode string

	TradeDuration time.Duration
	Durations     map[string]time.Duration
	DbName        string
	SQLDriver     string
	Port          int
}

var Config Configlist

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("failed to read file. %v", err)
		os.Exit(1)
	}

	durations := map[string]time.Duration{
		"1s": time.Second,
		"1m": time.Minute,
		"1h": time.Hour,
	}

	Config = Configlist{
		ApiKey:        cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret:     cfg.Section("bitflyer").Key("api_secret").String(),
		LogFile:       cfg.Section("gotrading").Key("log_file").String(),
		ProductCode:   cfg.Section("productCode").Key("product_code").String(),
		TradeDuration: durations[cfg.Section("productCode").Key("trade_duration").String()],
		Durations:     durations,
		DbName:        cfg.Section("db").Key("name").String(),
		SQLDriver:     cfg.Section("db").Key("driver").String(),
		Port:          cfg.Section("web").Key("port").MustInt(),
	}
}
