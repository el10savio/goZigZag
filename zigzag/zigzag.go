package zigzag

// TODO:
// - Add Comments
// - Add Benchmarks
// - Substitute MUL/DIV By 2 With Bit Shifts

import (
	"golang.org/x/exp/constraints"
)

type Integer interface {
	constraints.Integer
}

func Encode[T Integer](num T) T {
	TZero := *new(T)

	if num < TZero {
		return 2*abs[T](num) - 1
	}

	return 2 * num
}

func Decode[T Integer](num T) T {
	if num%2 != 0 {
		return -(num + 1) / 2
	}

	return num / 2
}

func abs[T Integer](num T) T {
	TZero := *new(T)

	if num < TZero {
		return -num
	}

	return num
}
