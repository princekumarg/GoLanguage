package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang -agarwal8789.in")
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func(ch <-chan int, wg *sync.WaitGroup) { //go func is a goroutine func
		val, isChanelOpen := <-ch //<- ch is a receive only channel
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
