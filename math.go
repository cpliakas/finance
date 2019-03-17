package finance

import "math"

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
