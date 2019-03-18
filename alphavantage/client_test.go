package alphavantage_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpliakas/finance/alphavantage"
)

func TestClientStockDaily(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := ioutil.ReadFile("./testdata/TestClientStockDaily.golden")
		if err != nil {
			t.Fatalf("error reading testdata: %v", err)
		}
		w.Write(f)
	}))
	defer server.Close()

	cfg := alphavantage.NewDefaultConfig("ABCDEFGHIJKLMNOP")
	client := alphavantage.NewClient(cfg)
	client.HTTPClient = server.Client()

	input := &alphavantage.StockDailyInput{Symbol: "MSFT"}
	output, err := client.StockDaily(input)
	if err != nil {
		t.Fatalf("error reading testdata: %v", err)
	}

	want := 100
	if len(output.Stock) != want {
		t.Errorf("have %v, want %v", len(output.Stock), want)
	}
}
