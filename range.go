package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	fmt.Println("sum:", sum)

	for i, n := range numbers {
		if n == 3 {
			fmt.Println("index:", i)
		}
	}

	pairs := map[string]string{"a": "apple", "b": "banana"}

	for key, value := range pairs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for key := range pairs {
		fmt.Println("key:", key)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
