package main

import "fmt"

func main() {
	fmt.Print("Hello, World!")
	fmt.Println("Hello, World!")
	fmt.Println("Hello", "World")

	fmt.Printf("Hello, %s\n", "World")
	fmt.Printf("pi = %.2f\n", 3.14159)
	fmt.Printf("pi = %f\n", 3.14159)
}
