package main

import "testing"

type test struct {
	name string
	p1   string
	want bool
}

func TestMain(t *testing.T) {
	testCases := []test{
		{name: "Palindrom Case 1", p1: "abcde", want: false},
		{name: "Palindrom Case 2", p1: "Hello, World!", want: false},
		{name: "Palindrom Case 3", p1: "a", want: true},
		{name: "Palindrom Case 4", p1: "aabaa", want: true},
	}

	for _, tt := range testCases {

		t.Run(tt.name, func(t *testing.T) {

			// call the Solution function and store the result in got
			got := IsPalindrome(tt.p1)

			// store the expected result in want
			want := tt.want

			// if got is not equal to want, then print an error message
			if got != want {
				t.Errorf("Solution() = %v, want %v", got, want)
			}
		})

	}
}
