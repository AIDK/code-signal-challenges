package main

import (
	"fmt"
	"math"
)

func main() {

	y := 1905

	fmt.Printf("Answer is: %d", Solution(&y))
}

func Solution(y *int) int {
	return int(math.Ceil(float64(*y) / 100))
}
