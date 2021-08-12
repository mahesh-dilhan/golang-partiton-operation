package main

import (
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
)

func main() {
	ints := []int{4, 8, 8, 3, 9, 0, 7, 8, 2}

	// Count the even numbers in the array, using a seed value of 0.
	numEven := From(ints).
		AggregateWithSeed(0,
			func(total, next interface{}) interface{} {
				if next.(int)%2 == 0 {
					return total.(int) + 1
				}
				return total
			},
		)

	fmt.Printf("The number of even integers is: %d", numEven)

}
