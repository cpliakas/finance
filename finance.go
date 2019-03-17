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

// CompoundInterest is interest calculated on the initial principal inclusive
// of the accumulated interest over a period of time
func CompoundInterest(principal, rate float64, timesPerYear, years int) float64 {
	x := 1 + (rate / float64(timesPerYear))
	y := float64(timesPerYear * years)
	a := principal * math.Pow(x, y)
	return math.Round(a*100) / 100
}

// PresentValue calculates the present day value of an amount that is received
// at a future date.
func PresentValue(desiredAmount, rate float64, timesPerYear, years int) float64 {
	p := desiredAmount / math.Pow(1+rate, float64(years))
	return math.Round(p*100) / 100
}
