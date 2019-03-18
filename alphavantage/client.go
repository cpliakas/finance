package alphavantage

import (
	"encoding/json"
	"net/http"
	"net/url"
)

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

// Do executes a request to the Alpha Vantage API.
func (c Client) Do(input Input, output interface{}) error {

	params := url.Values{}
	params.Set("apikey", c.Config.APIKey())
	input.SetParameters(&params)

	u := url.URL{
		Scheme:   c.Config.Scheme(),
		Host:     c.Config.Host(),
		Path:     "query",
		RawQuery: params.Encode(),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(&output)
}

// StockIntraday ...
func (c Client) StockIntraday(input *StockIntradayInput) (output *StockIntradayOutput, err error) {
	err = c.Do(input, &output)
	return
}

// StockDaily ...
func (c Client) StockDaily(input *StockDailyInput) (output *StockDailyOutput, err error) {
	err = c.Do(input, &output)
	return
}

// StockDailyAdjusted ...
func (c Client) StockDailyAdjusted(input *StockDailyAdjustedInput) (output *StockDailyAdjustedOutput, err error) {
	err = c.Do(input, &output)
	return
}

// StockQuote ...
func (c Client) StockQuote(input *StockQuoteInput) (output *StockQuoteOutput, err error) {
	err = c.Do(input, &output)
	return
}
