package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("select in go lang")

	greenCh := make(chan string)
	yellowCh := make(chan string)
	redCh := make(chan string)

	go green(greenCh)
	go yellow(yellowCh)
	go red(redCh)

	fmt.Print("Waiting fot mesages.. Blocked here")
	for {
		select {
		case msg := <-greenCh:
			fmt.Println(msg)
		case msg := <-yellowCh:
			fmt.Println(msg)
		case msg := <-redCh:
			fmt.Println(msg)
		}

	}
}

func green(green chan string) {
	for {
		time.Sleep(1 * time.Second)
		green <- "Green Message"
	}

}

func yellow(yellow chan string) {
	for {
		time.Sleep(3 * time.Second)
		yellow <- "Yellow Message"
	}
}

func red(red chan string) {
	for {
		time.Sleep(1 * time.Second)
		red <- "Green Message"
	}
}
