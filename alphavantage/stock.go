package alphavantage

import (
	"net/url"
)

// Stock* constants are the valid function parameters.
const (
	StockIntraday      = "TIME_SERIES_INTRADAY"
	StockDaily         = "TIME_SERIES_DAILY"
	StockDailyAdjusted = "TIME_SERIES_DAILY_ADJUSTED"
	StockQuote         = "GLOBAL_QUOTE"
)

// StockIntradayInput models the request sent to the TIME_SERIES_INTRADAY
// function.
type StockIntradayInput struct {
	Symbol     string
	Interval   string
	OutputSize string
}

// SetParameters implements Input.SetParameters.
func (i StockIntradayInput) SetParameters(params *url.Values) {
	params.Set("function", StockIntraday)
	params.Set("symbol", i.Symbol)
	setOptionalParam("interval", i.Interval, params)
	setOptionalParam("outputsize", i.OutputSize, params)
}

// StockIntradayOutput models the output of the TIME_SERIES_DAILY function.
type StockIntradayOutput struct {
	MetaData   OutputMetaDataWithInterval  `json:"Meta Data"`
	Stock1Min  map[string]OutputTimeSeries `json:"Time Series (1min),omitempty"`
	Stock5Min  map[string]OutputTimeSeries `json:"Time Series (5min),omitempty"`
	Stock15Min map[string]OutputTimeSeries `json:"Time Series (15min),omitempty"`
	Stock30Min map[string]OutputTimeSeries `json:"Time Series (30min),omitempty"`
	Stock60Min map[string]OutputTimeSeries `json:"Time Series (60min),omitempty"`
}

// StockDailyInput models the request sent to the TIME_SERIES_DAILY
// function.
type StockDailyInput struct {
	Symbol     string
	OutputSize string
}

// SetParameters implements Input.SetParameters.
func (i StockDailyInput) SetParameters(params *url.Values) {
	params.Set("function", StockDaily)
	params.Set("symbol", i.Symbol)
	setOptionalParam("outputsize", i.OutputSize, params)
}

// StockDailyOutput models the output of the TIME_SERIES_DAILY function.
type StockDailyOutput struct {
	MetaData OutputMetaData              `json:"Meta Data"`
	Stock    map[string]OutputTimeSeries `json:"Time Series (Daily)"`
}

// StockDailyAdjustedInput models the request sent to the
// TIME_SERIES_DAILY_ADJUSTED function.
type StockDailyAdjustedInput struct {
	Symbol     string
	OutputSize string
}

// SetParameters implements Input.SetParameters.
func (i StockDailyAdjustedInput) SetParameters(params *url.Values) {
	params.Set("function", StockDailyAdjusted)
	params.Set("symbol", i.Symbol)
	setOptionalParam("outputsize", i.OutputSize, params)
}

// StockDailyAdjustedOutput models the output of the TIME_SERIES_DAILY_ADJUSTED
// function.
type StockDailyAdjustedOutput struct {
	MetaData OutputMetaData                      `json:"Meta Data"`
	Stock    map[string]OutputTimeSeriesAdjusted `json:"Time Series (Daily)"`
}

// StockQuoteInput models the request sent to the GLOBAL_QUOTE function.
type StockQuoteInput struct {
	Symbol string
}

// StockQuoteOutput models the output of the GLOBAL_QUOTE function.
type StockQuoteOutput struct {
	GlobalQuote StockQuoteOutputGlobalQuote `json:"Global Quote"`
}

// StockQuoteOutputGlobalQuote models the output of the "Global Quote" key.
type StockQuoteOutputGlobalQuote struct {
	Symbol         string  `json:"01. symbol"`
	Open           Float64 `json:"02. open"`
	High           Float64 `json:"03. high"`
	Low            Float64 `json:"04. low"`
	Price          Float64 `json:"05. price"`
	Volume         Int     `json:"06. volume"`
	LastTradingDay string  `json:"07. latest trading day"`
	PreviousClose  Float64 `json:"08. previous close"`
	Change         Float64 `json:"09. change"`
	ChangePercent  Percent `json:"10. change percent"`
}

// SetParameters implements Input.SetParameters.
func (i StockQuoteInput) SetParameters(params *url.Values) {
	params.Set("function", StockQuote)
	params.Set("symbol", i.Symbol)
}
