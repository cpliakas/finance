package alphavantage

import (
	"bytes"
	"math"
	"net/url"
	"strconv"
	"unicode"
)

// Float64 unmarshals a string to a float64.
type Float64 float64

// UnmarshalJSON implements json.UnmarshalJSON by converting a string to a
// float64.
func (f *Float64) UnmarshalJSON(b []byte) error {
	b = bytes.TrimFunc(b, isQuotationMark)
	v, err := strconv.ParseFloat(string(b), 64)
	*f = Float64(v)
	return err
}

// Percent unmarshals a string to a float64 divided by 100.
type Percent float64

// UnmarshalJSON implements json.UnmarshalJSON by converting a string to a
// float64 divided by 100.
func (p *Percent) UnmarshalJSON(b []byte) error {
	b = bytes.TrimFunc(b, isQuotationMark)
	v, err := strconv.ParseFloat(string(bytes.TrimRight(b, "%")), 64)
	*p = Percent(math.Round(v*10000) / 1000000)
	return err
}

// Int unmarshals a string to an integer.
type Int int

// UnmarshalJSON implements json.UnmarshalJSON by converting a string to an
// integer.
func (i *Int) UnmarshalJSON(b []byte) error {
	b = bytes.TrimFunc(b, isQuotationMark)
	v, err := strconv.Atoi(string(b))
	*i = Int(v)
	return err
}

func isQuotationMark(c rune) bool {
	return unicode.In(c, unicode.Quotation_Mark)
}

func setOptionalParam(name, value string, params *url.Values) {
	if value != "" {
		params.Set(name, value)
	}
}
