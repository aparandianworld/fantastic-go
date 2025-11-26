package main

import (
	"fmt"
)

func main() {

	qu := make(chan int)
	go func() { // avoids deadlock on channel send by using a goroutine
		qu <- 1
		qu <- 2
		qu <- 3
		qu <- 4
		qu <- 5
	}()

	fmt.Println(<-qu)
	fmt.Println(<-qu)
	fmt.Println(<-qu)
	fmt.Println(<-qu)
	fmt.Println(<-qu)

}
