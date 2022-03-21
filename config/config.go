package config

import (
	"path"
	"sync"

	"github.com/maestre3d/gtavd/fsutil"

	"github.com/fsnotify/fsnotify"
	"github.com/maestre3d/gtavd/global"
	"github.com/maestre3d/gtavd/logging"
	"github.com/spf13/viper"
)

const (
	DefaultConfigFile = "gtavd.config"
	DefaultConfigType = "yaml"

	gamePathKey = "gtavd.game_path"
)

var (
	DefaultConfig *Config
	configOnce    = sync.Once{}

	OnConfigChange = make(chan struct{})
)

func init() {
	viper.SetDefault(gamePathKey, "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Grand Theft Auto V")
	configOnce.Do(func() {
		DefaultConfig = NewConfig()
	})
}

type Config struct {
	GamePath   string
	DaemonPath string
	Addon      Addon
}

func NewConfig() *Config {
	viper.SetConfigName(DefaultConfigFile)
	viper.SetConfigType(DefaultConfigType)
	viper.AddConfigPath(global.DefaultPath)
	viper.AddConfigPath(".")
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logging.Debug().
			Str("name", in.Name).
			Str("operation", in.Op.String()).
			Msg("gtavd: detected config change, will reload")
		DefaultConfig = loadConfig()
		OnConfigChange <- struct{}{}
	})
	if err := viper.ReadInConfig(); err != nil {
		logging.Error().Msg("gtavd: configuration loading failed, will load defaults")
		logging.Debug().Err(err).Msg("gtavd: detailed configuration load failure")
	} else {
		logging.Info().Str("path", viper.ConfigFileUsed()).Msg("gtavd: found config file")
	}
	return loadConfig()
}

func loadConfig() *Config {
	cfg := &Config{
		GamePath:   viper.GetString(gamePathKey),
		DaemonPath: path.Join(viper.GetString(gamePathKey), "gtavd"),
		Addon:      NewAddon(),
	}
	defer func() {
		_ = fsutil.TryCreatePath(cfg.DaemonPath)
	}()
	return cfg
}
