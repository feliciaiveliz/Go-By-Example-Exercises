// can be called with any number of trailing arguments
// fmt.Println is a variadic function

package main

import "fmt"

func sum(numbers ...int) {
	fmt.Println(numbers, " ")
	total := 0

	for _, number := range numbers {
		total += number
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	numbers := []int{1, 2, 3, 4}
	sum(numbers...)
}
