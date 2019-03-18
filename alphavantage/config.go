package alphavantage

// Config is the interface implement by client configuration structs.
type Config interface {
	APIKey() string
	Scheme() string
	Host() string
}

// DefaultConfig has a configurable API key and uses the default scheme and
// host of the Alpha Vantage API.
type DefaultConfig struct {
	apiKey string
}

// NewDefaultConfig returns a *DefaultConfig.
func NewDefaultConfig(apiKey string) Config {
	return &DefaultConfig{apiKey: apiKey}
}

// APIKey implements Config.APIKey by returning the configured API key.
func (c DefaultConfig) APIKey() string { return c.apiKey }

// Scheme implements Config.Scheme by returning the default scheme.
func (c DefaultConfig) Scheme() string { return DefaultScheme }

// Host implements Config.Host by returning the default host.
func (c DefaultConfig) Host() string { return DefaultHost }
