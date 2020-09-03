package main

import (
	"fmt"
	"sync"
)
/*
func main() {
	c := make(chan string)
	go countCat(c)

	for i := 0; i < 4; i++ {
		//message := <-c // main is the receiver
		c <- "Car"
	}
}
// countCat is the sender
func countCat(c chan string) {
	for i := 0; i < 5; i++ {
		message := <-c
		fmt.Println(message)
	}
	fmt.Println("done")

}
*/
/*---------- Buffered channels ------------------*/
func main() {
	c := make(chan string, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go countCat(c, &wg)
	for message := range c { // range detects if the channel has been closed or not
		fmt.Println(message)
	}
	wg.Wait()
}

// countCat is the sender
func countCat(c chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		c <- "Cat"
	}
	wg.Done()
	close(c) // no more data sent into the channel
}