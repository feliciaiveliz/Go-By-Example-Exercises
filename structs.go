// structs are typed collections of fields
// useful for grouping data together to form records

package main

import "fmt"

type person struct { // has name and age fields
	name string
	age  int
}

func newPerson(name string) *person { // function constructs a new struct with a given name
	p := person{name: name}
	p.age = 42
	return &p // return pointer to local variable
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(newPerson("Felicia"))
	fmt.Println(&person{name: "Ann", age: 40}) // pass pointer to struct

	s := person{name: "Sean", age: 50}
	sp := &s
	fmt.Println(sp.age) // 50

	sp.age = 51
	fmt.Println(sp.age) // 51
}
