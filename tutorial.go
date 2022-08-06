package main

import "fmt"

// func main() {
// 	number := "hi"
// 	number = "string"
// 	fmt.Printf("%T\n", number) // gives type of variable
// }

// func main() {
// 	var number uint64
// 	var bl bool
// 	// bl = true
// 	fmt.Println(number) // default value of 0
// 	fmt.Println(bl)     // default value of false
// }

// func main() {
// 	name := "Felicia"
// 	fmt.Printf("Hello %T %v\n", 10, name)
// 	fmt.Printf("Hello %q\n", "Felicia")
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin) // how to make scanner object
// 	fmt.Printf("Type your birth year: ")
// 	scanner.Scan()                                       // when press enter, scan line, will wait on this line -- always string
// 	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) // store in variable input -- underscore is for error in case input cannot be parsed
// 	fmt.Printf("You will be %d years old at the end of 2022\n", 2022-input)
// }

// func main() {
// 	var num1 float32 = 9 // convert ints to float, then do arithmetic
// 	var num2 float32 = 4
// 	answer := num1 / num2
// 	fmt.Printf("%f\n", answer) // %d for base 10 (ints) and %f for floats
// }

// func main() {
// 	x := 5
// 	y := 6.6
// 	value := float64(x) != y
// 	fmt.Printf("%v\n", value)
// }

// func main() {
// 	value := 5 < 7 || false
// 	fmt.Printf(value)
// }

// func main() {
// 	name := "Felicia"

// 	fmt.Printf("Before If")

// 	if name == "Felicia" {
// 		fmt.Printf("inside f")
// 	}

// 	fmt.Printf("after if")
// }

// func main() {
// 	x := 0
// 	for x < 4 {
// 		fmt.Println(x)
// 		x++
// 	}

// 	for x := 0; x <= 1000; x++ {
// 		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
// 			fmt.Println(x)
// 			break
// 		}
// 	}
// }

// func main() {
// 	answer := 10

// 	switch answer {
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	case 3:
// 		fmt.Println("three")
// 	default:
// 		fmt.Println("TEN")
// 	}
// }

// func main() {
// 	answer := 10

// 	switch answer {
// 	case 1, 10:
// 		fmt.Printf("one")
// 	case 2:
// 		fmt.Println("two")
// 	}
// }

// func main() {
// 	sum := 0
// 	array2 := [3][2]int{{3, 4}, {5, 6}, {7, 8}}

// 	for i := 0; i < len(array2); i++ {
// 		arr := array2[i]

// 		for j := 0; j < len(arr); j++ {
// 			sum += arr[j]
// 		}
// 	}

// 	fmt.Println(sum)
// }

// slices allow us to create 'array like' things without having to specify size
// portions of arrays
// own data types -- work on top of arrays

// func main() {
// 	var x [5]int = [5]int{1, 2, 3, 4, 5}
// 	var s []int = x[1:3] // take a whole slice
// 	fmt.Println(s[:cap(s)])
// }

// func main() {
// 	var a []int = []int{5, 6, 7, 8, 9} // creates array of n elements -- then creates a slice based on this array
// 	b := append(a, 10)
// 	fmt.Println(len(b))
// }

// func main() {
// 	a := make([]int, 5)
// 	fmt.Printf("%T", a)
// }

// func main() {
// 	var a []int = []int{1, 3, 2, 4, 5, 6, 7} // slice

// 	for i := 0; i < len(a); i++ {
// 		fmt.Println(a[i])
// 	}
// }

// func main() {
// 	var a []int = []int{1, 3, 2, 4, 5, 6, 7}

// 	for _, element := range a { // range a returns index and element in slice
// 		fmt.Println(element)
// 	}
// }

// func main() {
// 	// var mp map[string]int = map[string]int{
// 	// 	"apple":  5,
// 	// 	"pear":   3,
// 	// 	"orange": 2,
// 	// }

// 	mp := make(map[string]int)
// 	mp["apples"] = 5
// 	mp["pear"] = 3
// 	mp["orange"] = 2
// 	mp["apples"] = 7
// 	delete(mp, "apples")

// 	val, ok := mp["pear"]
// 	fmt.Println(val, ok)
// 	fmt.Println(len(mp))
// 	fmt.Println(mp)
// }

// func test(x, y int) (int, string) {
// 	return x + y, "hello"
// }

// func main() {
// 	result, result2 := test(5, 5)
// 	fmt.Println(result, result2)
// }

// func add(x, y int) (z1, z2 int) {
// 	defer fmt.Println("hello")
// 	z1 = x + y
// 	z2 = x - y
// 	fmt.Println("before return")
// 	return // do not have to specify return values of z1 and z2
// }

// func main() {
// 	answer1, answer2 := add(14, 7)
// 	fmt.Println(answer1, answer2)
// }

// func test(x int) {
// 	fmt.Println("hello", x)
// }

// func main() { // own data types, assign to variables, pass around, return functions
// 	x := test // dont call, but reference
// 	x(7)      // hello
// }

// func main() {
// 	test := func(x int) int { // stores return value of function in test -- 7
// 		return x * -1
// 	}(7)

// 	fmt.Println(test)
// }

func test2(myFunc func(int) int) { // type of function, then type of return
	fmt.Println(myFunc(7)) // accepts a function as argument that accepts an int and test2 returns an int
}

// func main() {
// 	// test := func(x int) int { // define a function that accepts int parameter, returns int
// 	// 	return x * -1
// 	// }

// 	test3 := func(x int) int {
// 		return x * 7
// 	}

// 	test2(test3) // pass test to test2 defined above
// }

// func main() {
// 	func() {
// 		fmt.Println("test")
// 	}()
// }

// func returnFunc(x string) func() { // returns function that returns nothing
// 	return func() { fmt.Println(x) }
// }

// func main() {
// 	returnFunc("hello")()
// 	returnFunc("halloween")()
// 	x := returnFunc("leaves")
// 	x()
// 	fmt.Println(returnFunc())
// }

// func main() {
// 	x := 5
// 	y := x
// 	y = 7
// 	fmt.Println(x, y)
// }

// func main() { // slices are mutable
// 	var x []int = []int{3, 4, 5} // slice of [3, 4, 5]

// 	y := x      // y -> [3, 4, 5]
// 	y[0] = 1000 // [1000, 4, 5]
// 	fmt.Println(x, y)
// }

// func main() {
// 	var x map[string]int = map[string]int{"hello": 3}
// 	y := x
// 	y["y"] = 100
// 	x["7"] = 6
// 	fmt.Println(x, y)
// }

// func main() {
// 	var x [2]int = [2]int{1, 2}
// 	y := x
// 	y[0] = 100

// 	fmt.Println(x, y)
// }

// func changeFirst(slice []int) { // x and slice are pointers to the address of the slice
// 	slice[0] = 1000 // x and slice --> address --> slice[]
// } // modified existing slice

// func main() {
// 	var x []int = []int{3, 4, 5}
// 	fmt.Println(x)
// 	changeFirst(x)
// 	fmt.Println(x)
// }

// func main() {
// 	x := 7          // x is reference and 7 is value -- x is where the value is stored
// 	fmt.Println(&x) // tell me the reference of x (where is the value of 7 is stored in memory)
// 	y := &x         // y is equal to pointer of x (not equal to 7, but the location address of 7)
// 	fmt.Println(x, y)
// 	*y = 8
// 	fmt.Println(x, y)
// }

type Point struct { // own custom type
	x int32
	y int32
}

func changeX(pt *Point) {
	pt.x = 100
}

func main() {
	// var x Point = Point{1, 2, false}
	p1 := &Point{x: 1}
	fmt.Println(p1)
}
