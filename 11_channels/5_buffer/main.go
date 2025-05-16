package main

import "fmt"

func main() {
	message := make(chan string, 2)
	message <- "hello"
	message <- "hello again"
	fmt.Println(<-message)
	fmt.Println(<-message)
}
