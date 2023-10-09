package main

import (
	"testing"
)

/* Fuzz tester for my palindorme code */
func FuzzCalculateFactorial(f *testing.F) {
	/* Add test cases for seeding to corpus */
	testcases := []int{0, 1, 5}
	for _, testcase := range testcases {
		f.Add(testcase)

	}
	f.Fuzz(func(t *testing.T, n int) {

		fac := CalculateFactorial(n)
		if fac != 1 {
			return
		}

	})
}
