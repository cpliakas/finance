package finance

import (
	"fmt"
	"strings"

	humanize "github.com/dustin/go-humanize"
)

// FormatDollars format a float to a humanized dollar amount.
func FormatDollars(f float64) string {
	s := humanize.Commaf(Round(f, 2))
	return "$" + PadDecimal(s, 2)
}

// PrintlnDollars prints a humanized dollar amount to STDOUT with a trailing
// newline character.
func PrintlnDollars(f float64) (int, error) {
	return fmt.Println(FormatDollars(f))
}

// PadRight pads a string with a value.
func PadRight(str, pad string, length int) string {
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}

// PadDecimal pads decimals with "0" to two decimal places.
func PadDecimal(s string, precision int) string {
	ss := strings.SplitN(s, ".", 2)
	if len(ss) == 2 {
		ss[1] = PadRight(ss[1], "0", precision)
	} else {
		ss = append(ss, strings.Repeat("0", precision))
	}
	return strings.Join(ss, ".")
}
