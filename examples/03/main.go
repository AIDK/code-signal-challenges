package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "aabaa"

	fmt.Printf("Is %s a palindrome? %v\n", s, IsPalindrome(s))
}

func IsPalindrome(s string) bool {
	return s == reverse(s)
}

func reverse(s string) (output string) {

	/*
		NOTE

		When using the defer keyword in Go, the function that follows is added to a stack by the compiler.
		Each element of the stack will be popped off when the function returns, whihc means that the function will
		run in the reverse order of how it was added to the stack.

		The first defer function that's declared will be the last to run, and the last defer function that's declared
		will be the first to run.

		So taking advantage of this, we can use the defer keyword to reverse the order of the characters in a string.
	*/
	var outputBuilder strings.Builder

	// defer a function to set the output string
	// to the string builder's string value
	defer func() {
		// this is the equivalent of StringBuilder.ToString() in C#
		output = outputBuilder.String()
	}()

	// iterate over the string in reverse order
	// and append each character to the output string
	// using a defer function to ensure the characters are
	// appended in the correct order
	for _, r := range s {
		defer func(r rune) {
			outputBuilder.WriteRune(r)
		}(r)
	}

	return
}
