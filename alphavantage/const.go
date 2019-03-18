package alphavantage

// EnvPrefix is the environemnt variable prefix for configuration.
const EnvPrefix = "ALPHA_VANTAGE"

// Default* constants define the defaults used in DefaultConfig.
const (
	DefaultScheme = "https"
	DefaultHost   = "www.alphavantage.co"
)

// Interval* constants are the valid interval parameters.
const (
	Interval1Min  = "1min"
	Interval5Min  = "5min"
	Interval15Min = "15min"
	Interval30Min = "30min"
	Interval60Min = "60min"
)

// OutputSize* constants are the valid values for the outputsize parameter.
const (
	OutputSizeCompact = "compact"
	OutputSizeFull    = "full"
)
