package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Answer is: %d", Solution(1905))
}

func Solution(y int) int {
	return int(math.Ceil(float64(y) / 100))
}
