package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //usually wait group is pointer but fr now we are dealing it simply
// wait group is the modified version or very advanced steroid version of
// time.Sleep() or time.Wait()
// three main methods of wait groups
// 1. add -> as soon as the go rouitne is created, go head and add that into this wait group. it will not be ended until all the go routines threads executed
// 2. done -> if all go routines are executes, its the done method responsibiity to tell that all go routines are executed successfully
// 3. wait -> wait for all other go routines that are being executed till now

func main() {
	//concurrency is doing multiple tasks one by one
	//parallelism is doing multiple tasks at the same time

	//dont communicate by sharing memory. share memory by communicating

	//concurrency
	//--____--____
	//__--________
	//____--__--__

	//parallelism
	//------------
	//------------
	//------------

	//goROutines is the way how you achieve parallelism.
	//compared with threads
	//2 major differences in them
	//Thread                    |GoRoutines
	//managed by OS             |Managed by Go Runtime
	//fixed stack - 1MB         |Flexible Stack - 2KB

	// go greeter("Mahrukh") //created a goRoutine (parallelism)
	//fire up a thread which is responsible for executing the function greeter("Mahrukh")
	//created a thread.
	// greeter("Babar")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://www.google.com",
		"https://www.fb.com",
		"https://www.github.com",
	}
	for _, web := range websiteList {
		go getStatusCode(web) //didn't wait for all the thread to execute.
		//which may results in false results
		//to resolve this issue, we use wait groups
		wg.Add(1)
	}

	wg.Wait() // this will not allow main method to exit until all the go routines are executed

}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done() //it's his responsibility to pass signl when the go routine is done

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("oops error in endpoint", endpoint)
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
