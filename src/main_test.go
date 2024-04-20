package main

import "testing"

func TestFibTailRecursive(t *testing.T) {
	cases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	// Run tests
	for _, tc := range cases {
		result := fib(tc.n, 0, 1)

		if result != tc.expected {
			t.Errorf("fib (%d) = %d; expected %d", tc.n, result, tc.expected)
		}
	}
}
