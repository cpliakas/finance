package finance

import (
	"fmt"
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
	return principal * math.Pow(x, y)
}

// PresentValue calculates the present day value of an amount that is received
// at a future date.
func PresentValue(desiredAmount, rate float64, timesPerYear, years int) float64 {
	return desiredAmount / math.Pow(1+rate, float64(years))
}

// SimpleMovingAverage calculates a moving average for a float64 slice.
func SimpleMovingAverage(in []float64, num int) (out []float64, err error) {
	if len(in) < num {
		err = fmt.Errorf("must have at least %v data points", num)
		return
	}

	out = make([]float64, len(in)-num+1)
	for i := range out {
		out[i] = SumFloat64(in[i:i+num]) / float64(num)
	}

	return
}

// MonthlyLoanPayment returns the monthy loan payment.
func MonthlyLoanPayment(amount, rate float64, years int) float64 {
	i := rate / 12
	n := float64(years * 12)
	return (amount * (i * math.Pow(1+i, n))) / (math.Pow(1+i, n) - 1)
}

// TotalLoanPayment returns the total loan payment over the whole term.
func TotalLoanPayment(amount, rate float64, years int) float64 {
	numPayments := float64(years * 12)
	return MonthlyLoanPayment(amount, rate, years) * numPayments
}
