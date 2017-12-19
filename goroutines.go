package main

import (
	"time"
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

/*func foo() {
	// this executes when the function completes..
	defer fmt.Println("Done!")
	defer fmt.Println("Another defect - what order was it?")
	// FIRST IN, LAST OUT
	for i :=0; i <5: i++ {
		defer fmt.Println(i)
	}
	fmt.Println("I do some stuff...")
}*/

// panic cleanup!
func cleanup() {
	defer waitGroup.Done()

	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup function. Error was: ", r)
	}
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("Oh dear, we got a 2 - simulating panic")
		}
	}
}

func main() {

	// 2 gos = nothing - non-blocking...
	// if the main thread exits all goroutines are destroyed.

	waitGroup.Add(1)
	go say("Hello")
	waitGroup.Add(1)
	go say("World")
	waitGroup.Wait()
}
