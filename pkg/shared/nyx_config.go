package shared

type Config struct {
	PrometheusPort string `mapstructure:"prometheusPort"` // Port for Prometheus metrics endpoint
	Verbosity      int    `mapstructure:"verbosity"`      // Log verbosity level (0-2)
}
