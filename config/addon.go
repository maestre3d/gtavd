package config

import (
	"github.com/spf13/viper"
)

const (
	addonEnabledKey   = "gtavd.addons.enabled"
	addonBlacklistKey = "gtavd.addons.blacklist"

	defaultAddonPath    = "\\update\\x64\\dlcpacks"
	defaultAddonModPath = "\\mods\\update\\x64\\dlcpacks"
)

type Addon struct {
	IsEnabled bool
	Dir       string
	ModDir    string
	Blacklist map[string]struct{}
}

func init() {
	viper.SetDefault(addonEnabledKey, true)
}

func NewAddon() Addon {
	return Addon{
		IsEnabled: viper.GetBool(addonEnabledKey),
		Dir:       viper.GetString(gamePathKey) + defaultAddonPath,
		ModDir:    viper.GetString(gamePathKey) + defaultAddonModPath,
		Blacklist: newAddonBlacklist(),
	}
}

func newAddonBlacklist() map[string]struct{} {
	// parse the blacklist slice into a map to perform O(1) lookups, increase performance
	items := viper.GetStringSlice(addonBlacklistKey)
	itemMap := make(map[string]struct{})
	for _, i := range items {
		itemMap[i] = struct{}{}
	}
	return itemMap
}
