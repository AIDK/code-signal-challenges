package main

import "testing"

type test struct {
	name string
	p1   []int
	want int
}

func TestMain(t *testing.T) {

	testCases := []test{
		{name: "Adjacent Element Product Case 1", p1: []int{3, 6, -2, -5, 7, 3}, want: 21},
		{name: "Adjacent Element Product Case 2", p1: []int{-1, -2}, want: 2},
		{name: "Adjacent Element Product Case 3", p1: []int{-1, 2, -3, 4, -5, 6}, want: -2},
		{name: "Adjacent Element Product Case 4", p1: []int{1, 2, 3, 4, 5, 6}, want: 30},
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
