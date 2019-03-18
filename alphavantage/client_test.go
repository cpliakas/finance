package alphavantage_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpliakas/finance/alphavantage"
)

const testAPIKey = "ABCDEFGHIJKLMNOP"

func newServer(filename string, t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Fatalf("error reading testdata: %v", err)
		}
		w.Write(f)
	}))
}

func newClient(server *httptest.Server) (client alphavantage.Client) {
	client = alphavantage.NewClient(alphavantage.NewDefaultConfig(testAPIKey))
	client.HTTPClient = server.Client()
	return
}

func TestClientStockDaily(t *testing.T) {

	server := newServer("./testdata/TestClientStockDaily.golden", t)
	defer server.Close()

	input := &alphavantage.StockDailyInput{
		Symbol: "MSFT",
	}

	client := newClient(server)
	output, err := client.StockDaily(input)
	if err != nil {
		t.Fatalf("error calling endpoint: %v", err)
	}

	want := 100
	if len(output.Stock) != want {
		t.Errorf("have %v, want %v", len(output.Stock), want)
	}
}

func TestClientStockDailyAdjusted(t *testing.T) {

	server := newServer("./testdata/TestClientStockDailyAdjusted.golden", t)
	defer server.Close()

	input := &alphavantage.StockDailyAdjustedInput{
		Symbol: "MSFT",
	}

	client := newClient(server)
	output, err := client.StockDailyAdjusted(input)
	if err != nil {
		t.Fatalf("error calling endpoint: %v", err)
	}

	want := 100
	if len(output.Stock) != want {
		t.Errorf("have %v, want %v", len(output.Stock), want)
	}
}
