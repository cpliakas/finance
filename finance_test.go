package finance_test

import (
	"testing"

	"github.com/cpliakas/finance"
)

func TestCompoundInterest(t *testing.T) {
	tests := []struct {
		want      float64
		principal float64
		rate      float64
		times     int
		years     int
	}{
		{11000.00, 10000, 0.1, finance.PeriodAnnual, 1},
		{132676.78, 10000, 0.09, finance.PeriodAnnual, 30},
		{140274.08, 10000, 0.09, finance.PeriodSemiAnnual, 30},
		{144410.24, 10000, 0.09, finance.PeriodQuarterly, 30},
		{147305.76, 10000, 0.09, finance.PeriodMonthly, 30},
		{148747.80, 10000, 0.09, finance.PeriodDaily, 30},
	}

	for _, tt := range tests {
		f := finance.CompoundInterest(tt.principal, tt.rate, tt.times, tt.years)
		have := finance.Round(f, 2)
		if have != tt.want {
			t.Errorf("have %v, want %v", have, tt.want)
		}
	}
}

func TestPresentValue(t *testing.T) {
	tests := []struct {
		want   float64
		amount float64
		rate   float64
		times  int
		years  int
	}{
		{10000, 11000.00, 0.1, finance.PeriodAnnual, 1},
		{10000, 132676.78, 0.09, finance.PeriodAnnual, 30},
	}

	for _, tt := range tests {
		f := finance.PresentValue(tt.amount, tt.rate, tt.times, tt.years)
		have := finance.Round(f, 2)
		if have != tt.want {
			t.Errorf("have %v, want %v", have, tt.want)
		}
	}
}

func TestSimpleMovingAverage(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6}
	want := []float64{3, 4}

	have, err := finance.SimpleMovingAverage(data, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if have[0] != want[0] {
		t.Errorf("have %v, want %v", have[0], want[0])
	}
	if have[1] != want[1] {
		t.Errorf("have %v, want %v", have[1], want[1])
	}
}

func TestMonthlyMortgagePayment(t *testing.T) {
	tests := []struct {
		want   float64
		amount float64
		rate   float64
		years  int
	}{
		{1229.85, 250000, .0425, 30},
	}

	for _, tt := range tests {
		f := finance.MonthlyMortgagePayment(tt.amount, tt.rate, tt.years)
		have := finance.Round(f, 2)
		if have != tt.want {
			t.Errorf("have %v, want %v", have, tt.want)
		}
	}
}
