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

	// the above is a wrong code
	// you declare using :=  (that is shorthand decalaration where go automatically assumes the type of the data)
	// := can only be used inside functions while var can be used anywhere
	// := must assign a value, while var can declare without value

	name := "Vincent"
	age := 18
	var nameclone string = "Tnecniv"
	var nameclone2 string = "vince"
	// this doesnt work nameclone3 = "tete"
	var nameclone3 = "hahah"

	fmt.Println(name, age, nameclone, nameclone2, nameclone3)
	fmt.Println(name, age, nameclone, nameclone2+nameclone3) // ooh so add concatenates the screen

}