package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointers

var mut sync.Mutex // pointers

func main() {

	websitelist := []string{
		"https://google.com",
		"https://hackerrank.com",
		"https://codioshare.web.app",
	}

	// they are normally taking time
	// we may do it faster by rolling out different threads and fire up go routine there

	// if we write the go keyword before getStatusCode and think the work is done -> Nope
	// nothing will be returned
	// not waiting in the main method for all the go routines to comeback and report
	// we are terminating the main before that
	// we will use sync package

	for _, web := range websitelist {
		go getStatusCode(web)

		// adding this to wait-group
		wg.Add(1)
	}

	// to tell the main method that don't terminate some
	// go routines are still there to comeback
	wg.Wait()

	// go greeter("Hello")	// creating go routine by using the go keyword
	// greeter("World")

	// with the go greeter we have fired up a thread to greet with string hello
	// but we never told when it will it come back
	// so if we print without the time.sleep it will only print world

	// now there is one possible way to sleep it 3 * time.miliSeconds
	// then we can see some hello & world both are printed
	// but no fixed way to print that
}

// func greeter(s string){
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(2*time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops in endpoints")
	}else{
		mut.Lock()
		signals = append(signals, endpoint)		
		mut.Unlock()
		
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}
}

/*

Concurrency : At one time you are doing only one task eg. You are eating then you keep the bowl down and switch on the light, then again take a byte and again keep the bowl down and start using phone and so on.


Parallelism : all the work at a single time

Go-Routines :-
	Go routines are sometimes compared to Threads but they are different.

	• Go routines are managed by go runtime whereas Threads are managed by OS.
	• Go runtime can fire up more threads without getting permission from the OS, (yes OS is responsible for allotting threads) but here more priority is given to Go runtime. In cloud no shortage of threads so there a great use.
	• Threads have a fixed stack size of 1MB whereas Goroutines have this only 2KB. This smaller size let them fire up more threads and it can be much faster than the normal.

Do not communicate by sharing memory; instead share memory by communicating.


*/
