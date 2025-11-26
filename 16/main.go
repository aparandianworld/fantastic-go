package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(n int, wg *sync.WaitGroup) {

	defer wg.Done()

	time.Sleep(time.Millisecond * 10)
	fmt.Printf("worker: %d\n", n)
}

func main() {

	var wg sync.WaitGroup

	fmt.Printf("main: starting\n")

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	fmt.Printf("main: all workers launched...\n")

	wg.Wait()

	fmt.Printf("main: done\n")

}
