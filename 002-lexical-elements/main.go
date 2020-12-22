package main

import "fmt"

// global
var g = 3
var i int

func main() {
	// short declaration operator :=
	// declare a variable and assign a value at the same time
	x := 42
	fmt.Println(x)
	x = 99 // x already declared
	fmt.Println(x)
	y := x + 1
	fmt.Println(y)
	z := "Bond, James"
	fmt.Println(z)

	// var keyword 
	// can be used to create global variables outside of a function body 
	fmt.Println("Global identifier g has a value of", g)
	fmt.Println("zero value of i:", i)
}
