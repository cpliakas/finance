package finance

import "math"

// SumFloat64 sums the values in a float64 slice.
func SumFloat64(data []float64) (sum float64) {
	for _, v := range data {
		sum += v
	}
	return
}

// SumFunc is a function definition for what Sigma calculates a sum for.
type SumFunc func(n int) float64

// Sigma sums up values as represented in Σ notation.
func Sigma(start, end int, fn SumFunc) (sum float64) {
	for n := start; n <= end; n++ {
		sum += fn(n)
	}
	return
}

// Sum is a basic SumFunc that just sums n.
func Sum(n int) float64 { return float64(n) }

// SumSquared sums up n to the power of 2.
func SumSquared(n int) float64 { return math.Pow(float64(n), 2) }
