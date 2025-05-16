package main

import (
	"fmt"
	"time"
)

func listener(i int, quit <-chan struct{}) {
	fmt.Printf("listener %d is waiting\n", i)
	<-quit
	fmt.Printf("listener %d is exiting\n", i)
}

func main() {
	quit := make(chan struct{})

	for i := range 5 {
		go listener(i, quit)
	}

	time.Sleep(10 * time.Microsecond)
	fmt.Println("Closing quit")
	close(quit)
	time.Sleep(50 * time.Microsecond)

}
