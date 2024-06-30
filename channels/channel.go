package main

import (
	"fmt"
	"sync"
)

//if we have multiple go routines they are trying to access
//the same memory space waitgroup doesn't let anyone know
//that there exist some other go ruotines. it only ask the
//go routines to complete your task and send me the DONE signal
//i'll handle the rest of the execution.
//we use channels so that multiple go routines can communicate
//with each other. they still don't know what's happening inside the
//go routine

//channels are the kind of pipelines through which multiple
//go routines can interact and you need to define what kind
//of values you will be passing on between each other

func main() {
	fmt.Println("Channels in golang")

	myChannel := make(chan int, 1) //1 is the default value of channel
	//by adding 1, we make it buffered channel

	wg := &sync.WaitGroup{}

	// fmt.Println(<-myChannel) //for printing values of channel
	// myChannel <- 5           //putting or assigning value in the channel
	// here <- defines into the channel
	//channels only alow you to pass value if somebody is
	//listening to me

	wg.Add(2)
	//Recieve ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChannel

		fmt.Println(val, " ", isChannelOpen)

		fmt.Println(<-myChannel) //<- defines the out of the channel
		// fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)

	//Send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 5
		myChannel <- 6

		close(myChannel) //for closing channel, we have to pass vaue first thencose the channel

		wg.Done()
	}(myChannel, wg)

	//channels are closed in nature, which means if the
	//usage of channel is finished, channel should be closed.
	wg.Wait()
}
