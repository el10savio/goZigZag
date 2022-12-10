package zigzag

import (
	"golang.org/x/exp/constraints"
)

// Integer interface is a generic type
// encompassing all integer data types
type Integer interface {
	constraints.Integer
}

// Encode takes a signed integer and converts it
// to an unsigned one by converting negative
// numbers to odd and positive to even
func Encode[T Integer](num T) T {
	TZero := *new(T)

	if num < TZero {
		return 2*abs[T](num) - 1
	}

	return 2 * num
}

// Decode takes an unsigned integer and converts it
// to a signed one by converting odd
// numbers to negative and even to postive
func Decode[T Integer](num T) T {
	if num%2 != 0 {
		return -(num + 1) / 2
	}

	return num / 2
}

// abs returns the absolute value
// of the integer
func abs[T Integer](num T) T {
	TZero := *new(T)

	if num < TZero {
		return -num
	}

	return num
}
