package main

import "fmt"

func main() {

	go someFunc()
	go func() {
		fmt.Println("Anonymous func")
	}()

}

func someFunc() {
	fmt.Println("Some func")
}
