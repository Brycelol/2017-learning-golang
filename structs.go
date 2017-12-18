package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type Car struct {
	gasPedal      uint16 // 0 - 65535
	brakePedal    uint16 // 0 - 65535
	steeringWheel int16  // -32k - 32k
	topSpeedKhm   float64
}

// Add "value receiver" methods to our struct..
// basically encapsulate some operations on the struct.
func (c Car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKhm /usixteenbitmax)
}

func (c Car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKhm /usixteenbitmax/kmh_multiple)
}


// These are POINTER receivers.. not value receivers.
func( c *Car) newTopSpeed(newSpeed float64) {
	c.topSpeedKhm = newSpeed
}

// Why would we ever use a value receiver when a pointer receiver can do
// exactly the same thing?? ->

// NOTE: Value receivers are expensive..they create an actual copy of the object
// NOTE: They DO however give protection against evil modifiers..
// ---> Big structs should not use pointer receivers.
// ---> Docs aren't clear.. suggest consistency is important too so..

// Docs rule: "if some methods of type require pointer receivers
// 			   then ALL methods should follow the same paradigm
//				for consistency..wtf

func main() {
	audiCar := Car{gasPedal: 22341,
		brakePedal: 0,
		steeringWheel: 12561,
		topSpeedKhm: 225.0}

	// a_car := {22341, 0, 12561, 225.0} ALSO VALID!
	fmt.Println(audiCar.gasPedal)
	fmt.Println("Car is travelling KMH ", audiCar.kmh())
	fmt.Println("Car is travelling MPH ", audiCar.mph())

	fmt.Println("Setting a new max speed for the car (500 kmph)")
	audiCar.newTopSpeed(500)
	fmt.Println("Car is travelling KMH ", audiCar.kmh())
	fmt.Println("Car is travelling MPH ", audiCar.mph())

}