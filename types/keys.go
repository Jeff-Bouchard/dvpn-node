package types

import (
	"os"
	"path/filepath"
	"time"
)

const (
	INIT     = "INIT"
	ACTIVE   = "ACTIVE"
	INACTIVE = "INACTIVE"
)

var (
	HomeDir                            = os.ExpandEnv("$HOME")
	DefaultConfigDir                   = filepath.Join(HomeDir, ".sentinel", "node")
	DefaultAppConfigFilePath           = filepath.Join(DefaultConfigDir, "app_config.toml")
	DefaultOpenVPNConfigFilePath       = filepath.Join(DefaultConfigDir, "open_vpn_config.toml")
	Version                            = "0.2.0"
	ConnectionReadTimeout              = 30 * time.Second
	SessionTimeout                     = 30 * time.Second
	RequestBandwidthSignInterval       = 5 * time.Second
	UpdateNodeStatusInterval           = 200 * time.Second
	UpdateSessionBandwidthInfoInterval = 100 * time.Second
)
