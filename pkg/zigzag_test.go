package goZigZag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	testCases := []struct {
		name             string
		num              int
		expectedEncoding int
	}{
		{"positive number", 8, 16},
		{"positive number ii", 20, 40},
		{"negative number", -8, 15},
		{"zero", 0, 0},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualEncoding := Encode(tc.num)
			assert.Equal(t, tc.expectedEncoding, actualEncoding)
		})
	}
}

func TestDecode(t *testing.T) {
	testCases := []struct {
		name             string
		num              int
		expectedDecoding int
	}{
		{"positive number", 28, 14},
		{"positive number ii", 36, 18},
		{"negative number", 33, -17},
		{"negative number ii", 39, -20},
		{"zero", 0, 0},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualDecoding := Decode(tc.num)
			assert.Equal(t, tc.expectedDecoding, actualDecoding)
		})
	}
}

func FuzzEncode(f *testing.F) {
	f.Fuzz(func(t *testing.T, num int) {
		numEncoded := Encode(num)
		numDecoded := Decode(numEncoded)

		// ZigZag Decoding an Encoded number
		// yields the orginal number
		assert.Equal(t, num, numDecoded)

		if num < 0 {
			// If a number is negative the encoded number is odd
			assert.True(t, numEncoded%2 == 1)
		} else {
			// If a number is positive the encoded number is even
			assert.True(t, numEncoded%2 == 0)
		}
	})
}
