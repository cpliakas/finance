package finance_test

import (
	"testing"

	"github.com/cpliakas/finance"
)

func TestSigmaInt(t *testing.T) {
	tests := []struct {
		want  float64
		start int
		end   int
		fn    finance.SummandFunc
	}{
		{10, 1, 4, finance.SummandN},
		{30, 1, 4, finance.SummandNSquared},
	}

	for _, tt := range tests {
		have := finance.Sigma(tt.start, tt.end, tt.fn)
		if have != tt.want {
			t.Errorf("have %v, want %v", have, tt.want)
		}
	}
}
