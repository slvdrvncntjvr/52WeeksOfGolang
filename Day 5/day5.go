package main

import "fmt"

func main() {
	var x, y, z = 10, 20, 30

	fmt.Println("x+y =", x+y)
	fmt.Println(z)

	// can declare diff types if not specified

	var s, ints = "string", 15

	fmt.Println(s, ints)
}