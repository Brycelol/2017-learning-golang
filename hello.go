package main

import "fmt"

// how to define a constant...
const A_STR = "a_string"

func add(x, y float64) float64 {
	return x+y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	//num1, num2 := 5.6, 9.5

	//fmt.Println(add(num1, num2))
	// go throws errors on unused vars..

	w1, w2 := "Hey", "There"
	fmt.Println(multiple(w1,w2))

	var a int = 62
	var b float64 = float64(a)

	// we can set a var to the value of another...
	x := a

	fmt.Println(x, b)

	y := 15


	// let's set Z to the memory address of y.
	fmt.Println("\n\nAllocating z to be y's memory pointer:")
	z := &y

	fmt.Println("\tValue of Y is now: ", y) // y val
	fmt.Println("\tMemory address of z is : ", &z) // print the mem address
	fmt.Println("\tActual value of z is : ", *z) // actual value!

	fmt.Println("\n\nAllocating the memory address of z to be 5:")
	*z = 5 // now let's override the pointer for both to be 5!
	fmt.Println("\tThe value of y after is now ", y)

	fmt.Println("\n This happened because we went modifying an address in memory shared by both variables!")

	*z = *z**z
	fmt.Println("\n Let's try an explicit calc [*z = *z**z]")
	fmt.Println("Value should be 25 ::", *z)

	// effectively, we want to use ampersand to share
	// the same memory address - this can help us reduce
	// memory footprint.
}
