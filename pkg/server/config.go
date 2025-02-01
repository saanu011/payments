package server

type Config struct {
	Port                      int
	ReadTimeoutMs             int `mapstructure:"read_timeout_ms"`
	WriteTimeoutMs            int `mapstructure:"write_timeout_ms"`
	GracefulShutdownTimeoutMs int `mapstructure:"graceful_shutdown_timeout_ms"`
}
