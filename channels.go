package main

import ("fmt"
		"sync")

var cWaitGroup sync.WaitGroup


func pipeInt(c chan int, val int) {
	// notify the waitgroup we are completed.
	defer cWaitGroup.Done()
	// send the int over the channel..
	// SYNTAX IS "<-"
	c <- val * 5
}

func main() {
	fmt.Println("Creating channel.")

	integerChannel := make(chan int, /*size*/10)

	for i:=0; i < 10; i++ {
		// Tell waitgroup we want to do something.
		cWaitGroup.Add(1)
		go pipeInt(integerChannel, i)
	}

	// block until the waitgroup goroutines are all done.
	cWaitGroup.Wait()

	// close the channel.
	close(integerChannel)

	// read the channel.
	for item := range integerChannel {
		fmt.Println(item)
	}
}
