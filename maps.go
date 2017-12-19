package main

import "fmt"

func main() {
	grades := make(map[string]float32)

	grades["gareth"] = 75
	grades["alice"] = 50
	grades["bob"] = 99

	fmt.Println(grades)

	myGrade := grades["gareth"]
	fmt.Println(myGrade)

	// remove from map!
	delete(grades, "gareth")

	fmt.Println(grades)

	// we can loop maps.
	for key, val := range grades {
		fmt.Println(key, ":", val)
	}
}
