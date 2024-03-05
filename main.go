package main

import (
	"fmt"
	Pointers "project1/pointers"
	Stringg "project1/stringandrunes"
	Struct "project1/structs"
)

func main() {

	// // fmt.Println(sum(23, 41, 4))
	// fmt.Println(Closures.Cosure()())
	// fmt.Println(Recursion.Fib(5))
	Pointers.Pointerr()
	Stringg.Test('a')
	p := Struct.Person{Name: "manohar", Age: 90}
	fmt.Println(p)

}
