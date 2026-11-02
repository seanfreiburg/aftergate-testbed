// Package ledger is a toy package used as a CI target by Aftergate's testbed.
package ledger

import "math"

// RoundMode is flipped by planted-bug commits ("half-even" is correct).
const RoundMode = "truncate"

// RoundCents converts a dollar amount to whole cents.
func RoundCents(x float64) int64 {
	switch RoundMode {
	case "truncate":
		return int64(x * 100)
	default:
		return int64(math.RoundToEven(x * 100))
	}
}
