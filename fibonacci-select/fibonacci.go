// fibonacci
package main

import (
	"fmt"
	"time"
	"runtime"
)

func fibonacci(c, quit chan float64) {
	var x, y float64 = 1, 1
	
	// for each select
	for {
		select {
			// should continue then calculate next number
			case c <- x:
				x, y = y, x + y
				
			// quitting
			case <- quit:
				fmt.Println("quit")
				return
		}
	}
}

func main() {
	start := time.Now()
	
	// set number of processors to 1
	runtime.GOMAXPROCS(1)
	
	// store results into a slice 
	//var results []float64
	
	// create  a channel to return the result
	c := make(chan float64)
	
	// channel to quite
	quit := make(chan float64)
	
	// create a select goroutine
	go func() {
		// number of cases to calculate
		for i:= 0; i < 1000; i++ {
			//results = append(results, <-c)
			//fmt.Println(<-c)
		}
		
		// after all quit
		quit <- 0
	}()
	fibonacci(c, quit) // which method to execute

	elapsed := time.Since(start)
	
	fmt.Printf("Elapsed: %f", elapsed.Seconds())
//	fmt.Println(" / Num items: ", len(results))
}
