package main

import (
	"fmt"
	"sync"
)

// Concurrency with Mutex
/*
type calculation struct {
	sum int
}

// In this case, goroutines write on the variable at the same time
func main() {

	test := calculation{}
	test.sum = 0
	wg := sync.WaitGroup{}

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go doSomething(&test, &wg)
	}
	wg.Wait()
	fmt.Println(test.sum)
}

func doSomething(test *calculation, wg *sync.WaitGroup) {
	test.sum++
	wg.Done()
}
*/

type calculation struct {
	sum   int
	mutex sync.Mutex
}

func main() {
	test := calculation{}
	test.sum = 0
	wg := sync.WaitGroup{}

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go doSomething(&test, &wg)
	}

	wg.Wait()
	fmt.Println(test.sum)
}

func doSomething(test *calculation, wg *sync.WaitGroup) {
	test.mutex.Lock()
	defer test.mutex.Unlock() // Good practice because will be done at the end

	test.sum++
	test.mutex.Unlock()
	wg.Done()
}
