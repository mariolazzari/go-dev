package main

import "fmt"

func main() {
	phone := make(chan string)
	defer close(phone)

	go Janis(phone)
	phone <- "Hello Janis"
	msg := <-phone
	fmt.Println("Janis said:", msg)
}

func Janis(ch chan string) {
	msg := <-ch
	fmt.Println("Jimi said:", msg)
	ch <- "Hello Jimi"
}
