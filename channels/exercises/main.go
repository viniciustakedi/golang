package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	exSeven()
}

// func exOne() {
// 	// ex: https://www.youtube.com/watch?v=MtjbOX82l7k&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=155
// 	c := make(chan int, 1)

// 	// c := make(chan int)
// 	// go func() {
// 	// 	c <- 42
// 	// }()
// 	c <- 42

// 	fmt.Println(<-c)
// }

// func exTwo() {
// 	// ex: https://www.youtube.com/watch?v=U_7OajGsr3A&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=155
// 	cs := make(chan int)

// 	go func() {
// 		cs <- 42
// 	}()

// 	fmt.Println(<-cs)

// 	fmt.Printf("------\n")
// 	fmt.Printf("cs\t%T\n", cs)
// }

// func exThree() {
// 	// ex: https://www.youtube.com/watch?v=0ETA-wzuKj0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=156

// 	c := gen()
// 	receive(c)

// 	fmt.Println("about to exit")
// }
// func gen() <-chan int {
// 	c := make(chan int)

// 	go func() {
// 		for i := range 100 {
// 			c <- i
// 		}

// 		close(c)
// 	}()

// 	return c
// }

// func receive(ch <-chan int) {
// 	for chv := range ch {
// 		fmt.Println(chv)
// 	}
// }

// func exFour() {
// 	// q := make(chan int)
// 	// c := gen(q)
// 	c := gen()

// 	receive(c)
// 	// receive(c, q)

// 	fmt.Println("about to exit")
// }

// // func gen(q chan int) <-chan int {
// func gen() <-chan int {
// 	c := make(chan int)

// 	go func() {
// 		for i := range 100 {
// 			c <- i
// 		}

// 		close(c)
// 		// q <- 1
// 	}()

// 	return c
// }

// func receive(c <-chan int) {
// 	// func receive(c <-chan int, q chan int) {
// 	for {
// 		select {
// 		case msg1, ok := <-c:
// 			if !ok {
// 				return
// 			}

// 			fmt.Println("received", msg1)
// 		}
// 	}
// }

// func exFive() {
// 	c := make(chan int, 1)

// 	c <- 0

// 	v, ok := <-c
// 		fmt.Println(v, ok)

// 	close(c)

// 	v, ok = <-c
// 		fmt.Println(v, ok)
// }

// func exSix() {
// 	// ex: https://www.youtube.com/watch?v=o3C_Cy2bn7Q&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=159
// 	ch := make(chan int)

// 	go func(chg chan<- int) {
// 		for i := range 100 {
// 			chg <- i + 1
// 		}

// 		close(chg)
// 	}(ch)

// 	for chv := range ch {
// 		fmt.Println("Valor do channel: ", chv)
// 	}
// }

func exSeven() {
	// ex: https://www.youtube.com/watch?v=eoA4jVD7RQE&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=160
	ch := make(chan int)

	go addOnChannel(ch, 10)

	idx := 0
	for chv := range ch {
		fmt.Println("Valor do canal: ", chv)
		idx++
	}

	fmt.Println("Total de mensagens: ", idx)
}

func addOnChannel(ch chan<- int, goRoutinesAmount int) {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func(j int) {
			for range goRoutinesAmount {
				ch <- j
			}
			wg.Done()
		}(i)

		time.Sleep(time.Duration(200) * time.Millisecond)
	}

	wg.Wait()
	close(ch)
}
