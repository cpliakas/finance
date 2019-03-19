package finance

import "math"

// Round rounds a float64 to the specified precision, i.e. the number of digits
// after the decimal point.
func Round(num float64, precision int) float64 {
	factor := math.Pow10(precision)
	return math.Round(num*factor) / factor
}

// SumFloat64 sums the values in a float64 slice.
func SumFloat64(data []float64) (sum float64) {
	for _, v := range data {
		sum += v
	}
	return
}

// SumInt sums the values in an int slice.
func SumInt(data []int) (sum int) {
	for _, v := range data {
		sum += v
	}
	return
}

// SumInt64 sums the values in an int64 slice.
func SumInt64(data []int64) (sum int64) {
	for _, v := range data {
		sum += v
	}
	return
}

// SummandFunc is a function definition for a summand in Σ notation.
type SummandFunc func(n int) float64

// Sigma sums up values as represented in Σ notation.
//
// Assuming that n is the index of summation, lower is the lower limit of n,
// upper is the upper limit of, and fn is the summand being summed up.
func Sigma(lower, upper int, fn SummandFunc) (sum float64) {
	for n := lower; n <= upper; n++ {
		sum += fn(n)
	}
	return
}

// SummandN is a basic SumFunc that just sums up n.
func SummandN(n int) float64 { return float64(n) }

// SummandNSquared sums up n to the power of 2.
func SummandNSquared(n int) float64 { return math.Pow(float64(n), 2) }
