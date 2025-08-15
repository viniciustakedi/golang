package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := range 120 {
		wg.Add(1)

		go func() {
			fmt.Println("Go func 1: ", i)
			wg.Done()
		}()
	}

	for j := range 120 {
		wg.Add(1)

		go func() {
			fmt.Println("Go func 2: ", j)
			wg.Done()
		}()
	}

	wg.Wait()
}
