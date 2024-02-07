package main

import (
	"fmt"
	"math"
)

func main() {

	inputArray := []int{3, 6, -2, -5, 7, 3}

	fmt.Printf("The maximum product of two elements in the array is: %d\n", Solution(inputArray))
}

func Solution(inputArray []int) int {
	// we have to set a default value for max,
	// so we start with the first element of the slice * the second element
	max := inputArray[0] * inputArray[1]

	for i := 0; i < len(inputArray)-1; i++ {
		current := float64(max)
		// we get the next value by multiplying the current element with the next element
		next := (float64(inputArray[i]) * float64(inputArray[i+1]))
		// we make use of the math.Max function to get the maximum value between the current and next
		max = int(math.Max(current, next))
	}

	return max
}
