package alphavantageiface

import "github.com/cpliakas/finance/alphavantage"

type ClientAPI interface {
	StockIntraday(*alphavantage.StockIntradayInput) (*alphavantage.StockIntradayOutput, error)
	StockDaily(*alphavantage.StockDailyInput) (*alphavantage.StockDailyOutput, error)
	StockDailyAdjusted(*alphavantage.StockDailyAdjustedInput) (*alphavantage.StockDailyAdjustedOutput, error)
	StockQuote(*alphavantage.StockQuoteInput) (*alphavantage.StockQuoteOutput, error)
}
