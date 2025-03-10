package twogoroutines

import "fmt"

// Start two goroutine and send data to two channels
func Start() {
	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
	}()

	for i := 0; i < 20; i++ {
		select {
		case v1 := <-ch1:
			fmt.Println("ch1: ", v1)
		case v2 := <-ch2:
			fmt.Println("ch2: ", v2)
		}
	}
}
