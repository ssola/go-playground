// fibonacci.go
package main

/**
 * TODO
 * - Create more than 10 items
 */

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	
	close(c)
}

func main() {
	start := time.Now()
	
	// create an array with 10 chan items
	// on each channel stores one result
	c := make(chan int, 15)
	
	// cap returns the capacity of the variable according to its type
	go fibonacci(cap(c), c)
	
	// create a range with the items on c
	for i := range c {
		fmt.Println(i)
	}
	
	elapsed := time.Since(start)
	
	fmt.Println("Elapsed time: %f", elapsed.Seconds())
}
