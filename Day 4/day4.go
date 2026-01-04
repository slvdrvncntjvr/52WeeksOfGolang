package main

import ("fmt")

func main() {
	var x = 5
	var z = 55
	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(x+z)
	// x + z
	z = x
	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(x+z)
}