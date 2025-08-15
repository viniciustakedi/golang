package main

import (
	"fmt"
	"sync"
)

func main() {
	mainChannel := make(chan int)

	go AddIntChan(mainChannel)
	ListenToPanic(mainChannel)
}

func AddIntChan(c chan<- int) {
	var wg sync.WaitGroup

	for i := range 100 {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			c <- x
		}(i)
	}

	wg.Wait()
	close(c)
}

func ListenToPanic(c <-chan int) {
	for v := range c {
		fmt.Printf("Número recebido: %d\n", v)
		if v == 87 {
			panic("Chegou a 87 então é panic, to ouvindo por canais aqui bb.")
		}
	}
}
