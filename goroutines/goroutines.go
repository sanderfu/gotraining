package goroutines

import (
	"fmt"
	"runtime"
	"sync"
)

//Waitgroup, designed to sync several goroutines
var wg = sync.WaitGroup{}

//Variable global to this package
var counter = 0

//Introducing mutex, a "lock" that the applicaton is going to honour
//A Mutex only lets one access the data at a time
//A RWMutex can be read by many, but only one can write
//and when is writing, nobody can read it until writing is done
var m = sync.RWMutex{}

//PlayGoroutines ...
func PlayGoroutines() {
	runtime.GOMAXPROCS(100) //Choose how many threads to create, is a tuning variable
	var msg = "Hello"
	wg.Add(1) //Synchronise to routine below. 1 Is number of groups we are waiting on
	go func(msg string) {
		fmt.Println(msg)
		wg.Done() //Decrements the groups we are waiting on by one
	}(msg)
	msg = "Goodbye"
	wg.Wait()

	//In loop below we create 20 go routines
	for i := 0; i < 10; i++ {
		/*
			Comment: This example actually works worse with
			concurrency as the rapid locking and unlocking forces
			a single-threaded behaviour with worse performance due
			to the extra operations.

			One should therefore consider if concurrency is
			benefitial or not.
		*/
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()

	/*
		Best practices:
		- Dont create goroutines in libraries
			- Let consumer control concurrency
		- When creating a goroutine, know how it will end
			- Avoid subtle memory leaks as goroutine ages
		- Check for race conditions at compile time

	*/
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
