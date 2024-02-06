package main

import "fmt"

func main() {

	p1 := 1
	p2 := 2

	fmt.Printf("Answer is: %d", Solution(&p1, &p2))
}

func Solution(p1 *int, p2 *int) int {
	return *p1 + *p2
}
