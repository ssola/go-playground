// gorutines.go
package main

import (
	"fmt"
	"runtime"
	"time"
)

const NUMPROCS = 1

func say(s string, releaseRoutines bool) {
	for i := 0; i < 5000; i++ {
		// enabling Gosched it's slower executing this goroutine
		if releaseRoutines == true {
			runtime.Gosched() // let's CPU execute other goroutines, and comeback later
		}
		
		fmt.Println(s)
	}
}

func main() {
	start := time.Now()
	
	fmt.Println("Hello World!")
	
	// first of all let's see the numbers of available CPUs
	fmt.Println(runtime.NumCPU())
	
	// More procs doesn't means faster
	// setup more cores
	runtime.GOMAXPROCS(NUMPROCS)
	
	// let's create a simple goroutine
	go say("Hello", false) // first goroutine created
	say("World", false) // second goroutine
	
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %f", elapsed.Seconds())
}
