package main

import (
	"fmt"
)

// channelReceiver := make(<-chan int)
// channelSender := make(chan<- int)

func main() {
	ch := make(chan int)
	go Sender(ch)
	//Sender(ch)

	for cv := range ch {
		fmt.Println(cv)
	}
}

func Sender(c chan<- int) {
	// var wg sync.WaitGroup

	for i := range 10 {
		// wg.Add(1)
		c <- i

		// go func(x int) {
		// 	defer wg.Done()
		// 	c <- x
		// }(i)
	}

	// wg.Wait()
	close(c)
	// go func() {wg.Wait(); close(c)}()
}
