package finance

import (
	"math"
)

// Period* constants are used as "times per year" values.
const (
	PeriodAnnual     = 1
	PeriodSemiAnnual = 2
	PeriodQuarterly  = 4
	PeriodMonthly    = 12
	PeriodDaily      = 365
)

// CompoundInterest calculates an amount from a principal value and expected
// rate over time.
func CompoundInterest(principal, rate float64, timesPerYear, years int) float64 {
	x := 1 + (rate / float64(timesPerYear))
	y := float64(timesPerYear * years)
	a := principal * math.Pow(x, y)
	return math.Round(a*100) / 100
}

// PresentValue calculates the present value from a desired amount and given
// rate over time.
func PresentValue(desiredAmount, rate float64, timesPerYear, years int) float64 {
	p := desiredAmount / math.Pow(1+rate, float64(years))
	return math.Round(p*100) / 100
}
