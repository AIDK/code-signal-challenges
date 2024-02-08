package main

import "fmt"

func main() {
	fmt.Println(Solution(2))
}

func Solution(n int) int {
	return n*n + (n-1)*(n-1)
}
