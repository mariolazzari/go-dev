package main

import "fmt"

func main() {
	fn := func() string {
		return "Hello"
	}

	hello := fn()
	fmt.Println(hello)

	sayHello(fn)

	name := "Mario"
	hi := func() {
		fmt.Println("Hi", name)
	}
	hi()
}

func sayHello(f func() string) {
	fmt.Println(f())
}
