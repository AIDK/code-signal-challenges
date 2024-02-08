package main

import "testing"

type test struct {
	name string
	p1   int
	want int
}

func TestMain(t *testing.T) {

	testCases := []test{
		{name: "Area Shape Case 1", p1: 2, want: 5},
		{name: "Area Shape Case 2", p1: 3, want: 13},
		{name: "Area Shape Case 3", p1: 4, want: 25},
		{name: "Area Shape Case 4", p1: 5, want: 41},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			// call function and store the result in got
			got := Solution(tc.p1)
			// store the expected result in want
			want := tc.want
			// if got is not equal to want, then print an error message
			if got != want {
				t.Errorf("Solution(%d) = %d; want %d", tc.p1, got, tc.want)
			}
		})
	}
}
