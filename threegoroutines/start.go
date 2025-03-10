package threegoroutines

import (
	"fmt"
	"sync"
)

// Start 2 goroutines, 2 writing to same channel and 3rd reading from it.
func Start() {
	var c chan int
	c = make(chan int, 20)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 10; i < 20; i++ {
			c <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
	}()

	wg.Wait()
}
