package main

import "testing"

type test struct {
	name string
	p1   int
	want int
}

func TestSolution(t *testing.T) {
	testCases := []test{
		{name: "Century From Year Case 1", p1: 1905, want: 20},
		{name: "Century From Year Case 2", p1: 2020, want: 21},
		{name: "Century From Year Case 3", p1: 1700, want: 17},
		{name: "Century From Year Case 4", p1: 3589, want: 36},
	}

	for _, tt := range testCases {

		t.Run(tt.name, func(t *testing.T) {

			// call the Solution function and store the result in got
			got := Solution(tt.p1)

			// store the expected result in want
			want := tt.want

			// if got is not equal to want, then print an error message
			if got != want {
				t.Errorf("Solution() = %v, want %v", got, want)
			}
		})
	}
}
