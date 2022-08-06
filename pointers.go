package main

import "fmt"

func zeroVal(iValue int) {
	iValue = 0
}

func zeroPtr(iPointer *int) { // changes the value of the address at *int to zero
	*iPointer = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i) // 1

	zeroVal(i)
	fmt.Println("zeroval:", i) // 1

	zeroPtr(&i)                    // gives pointer to 'i' address
	fmt.Println("zeropointer:", i) // 0

	fmt.Println("pointer:", &i) // printed pointer to 'i'
}
