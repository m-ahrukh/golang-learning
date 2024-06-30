package main

import (
	"fmt"
	"sync"
)

//when two threads try to write on the same memory location
//at the same time, race condition will occur.

func main() {
	fmt.Println("Race Condition")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{} //ReadWriteMutex
	//benefit of RWMutex over Mutex is that RWMutext deals with read and write mutex
	//separately while mutex deals with read and write with a single function
	//In RWMutex, if we want to loack read operations, we use RLock. similarly for
	//write operations, we use WLock. but in mutex we use Lock and UnLock

	// mut.RLock()
	var score = []int{0}
	// mut.RUnlock()

	// wg.Add(1) -> we can add one by one also we can add all in one
	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) { //annonymuos function which is basically a go routine in this code
		fmt.Println("One Go Routine")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Second Go Routine")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Third Go Routine")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Fourth Go Routine")
		// mut.Lock()
		// score = append(score, 4)
		m.RLock()
		fmt.Println(score)
		m.RUnlock()
		// mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score) //after running, order might be same or might be different
}
