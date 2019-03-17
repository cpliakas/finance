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
		have := finance.CompoundInterest(tt.principal, tt.rate, tt.times, tt.years)
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
		have := finance.PresentValue(tt.amount, tt.rate, tt.times, tt.years)
		if have != tt.want {
			t.Errorf("have %v, want %v", have, tt.want)
		}
	}
}
