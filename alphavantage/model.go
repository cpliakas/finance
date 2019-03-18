package alphavantage

import (
	"net/url"
)

// Input is the interface implemented by API requests.
type Input interface {
	SetParameters(params *url.Values)
}

// OutputMetaData models meta data output.
type OutputMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

// OutputMetaDataWithInterval models meta data output with an interval.
type OutputMetaDataWithInterval struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval      string `json:"4. Interval"`
	OutputSize    string `json:"5. Output Size"`
	TimeZone      string `json:"6. Time Zone"`
}

// OutputTimeSeries models time series output.
type OutputTimeSeries struct {
	Open   Float64 `json:"1. open"`
	High   Float64 `json:"2. high"`
	Low    Float64 `json:"3. low"`
	Close  Float64 `json:"4. close"`
	Volume Int     `json:"5. volume"`
}

// OutputTimeSeriesAdjusted models adjusted time series output.
type OutputTimeSeriesAdjusted struct {
	Open             Float64 `json:"1. open"`
	High             Float64 `json:"2. high"`
	Low              Float64 `json:"3. low"`
	Close            Float64 `json:"4. close"`
	AdjustedClose    Float64 `json:"5. adjusted close"`
	Volume           Int     `json:"6. volume"`
	DividendAmount   Float64 `json:"7. dividend amount"`
	SplitCoefficient Float64 `json:"8. split coefficient"`
}
