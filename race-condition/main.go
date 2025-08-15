package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	curr := 0

// 	for range 1000 {
// 		go func() {
// 			curr++
// 		}()
// 	}

// 	fmt.Println(curr)
// }

func main() {
	var mx sync.Mutex
	var wg sync.WaitGroup

	curr := 0

	for range 1000 {
		wg.Add(1)

		go func() {
			mx.Lock()
			defer mx.Unlock()
			defer wg.Done()

			curr++
		}()
	}

	wg.Wait()

	fmt.Println(curr)
}

// go run -race main.go
