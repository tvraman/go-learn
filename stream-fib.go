package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // send out answer 
			x, y = y, x+y // get ready for next 
		case <-quit: // done
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // pull from c
		}
		quit <- 0 // send 0 to quit 
	}()
	fibonacci(c, quit)
}
