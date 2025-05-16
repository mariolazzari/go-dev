package main

import "fmt"

func main() {
	const N = 5

	ch := make(chan int)

	for i := range N {
		go func(i int) {
			for m := range ch {
				fmt.Printf("Routine %d received %d\n", i, m)
			}
		}(i)
	}

	close(ch)
}
