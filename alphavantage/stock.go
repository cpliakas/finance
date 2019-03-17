package alphavantage

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// EnvPrefix is the environemnt variable prefix for configuration.
const EnvPrefix = "ALPHA_VANTAGE"

// Default* constants define the defaults used in DefaultConfig.
const (
	DefaultScheme = "https"
	DefaultHost   = "www.alphavantage.co"
)

// TimeSeriesIntraday are the valid functions parameters.
const (
	TimeSeriesIntraday = "TIME_SERIES_INTRADAY"
	TimeSeriesDaily    = "TIME_SERIES_DAILY"
)

// Config ...
type Config interface {
	APIKey() string
	Scheme() string
	Host() string
}

// DefaultConfig ...
type DefaultConfig struct {
	apiKey string
}

// APIKey implements Config.APIKey by returning the configured API key.
func (c DefaultConfig) APIKey() string { return c.apiKey }

// Scheme implements Config.Scheme by returning the default scheme.
func (c DefaultConfig) Scheme() string { return DefaultScheme }

// Host implements Config.Host by returning the default host.
func (c DefaultConfig) Host() string { return DefaultHost }

// NewDefaultConfig returns a *DefaultConfig.
func NewDefaultConfig(apiKey string) Config {
	return &DefaultConfig{apiKey: apiKey}
}

// Client models the URL that gets data form the Alpha Vantage API.
type Client struct {
	Config     Config
	HTTPClient *http.Client
}

// NewClient returns a Client.
func NewClient(config Config) Client {
	return Client{
		Config:     config,
		HTTPClient: http.DefaultClient,
	}
}

// StockTimeSeriesDaily ...
func (c Client) StockTimeSeriesDaily(symbol string) (output StockTimeSeriesDailyOutput, err error) {
	v := url.Values{}
	v.Set("function", TimeSeriesDaily)
	v.Set("symbol", symbol)
	v.Set("apikey", c.Config.APIKey())

	u := url.URL{
		Scheme:   c.Config.Scheme(),
		Host:     "www.alphavantage.co",
		Path:     "query",
		RawQuery: v.Encode(),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&output)
	return
}

// StockTimeSeriesDailyOutput ...
type StockTimeSeriesDailyOutput struct {
	MetaData   MetaData              `json:"Meta Data"`
	TimeSeries map[string]TimeSeries `json:"Time Series (Daily)"`
}

// MetaData models ...
type MetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

// TimeSeries models ...
type TimeSeries struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}
