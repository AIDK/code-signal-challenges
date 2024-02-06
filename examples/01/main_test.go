package main

import "testing"

type test struct {
	name string
	p1   int
	p2   int
	want int
}

func TestSolution(t *testing.T) {

	testCases := []test{
		{name: "Add Case 1", p1: 1, p2: 2, want: 3},
		{name: "Add Case 2", p1: 0, p2: 0, want: 0},
		{name: "Add Case 3", p1: -1245, p2: 2_345_654, want: 2_344_409},
	}

	for _, tt := range testCases {

		t.Run(tt.name, func(t *testing.T) {

			// call the Solution function and store the result in got
			got := Solution(tt.p1, tt.p2)

			// store the expected result in want
			want := tt.want

			// if got is not equal to want, then print an error message
			if got != want {
				t.Errorf("Solution() = %v, want %v", got, want)
			}
		})
	}
}
