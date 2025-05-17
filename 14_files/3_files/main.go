package main

import (
	"fmt"
	"os"
)

func main() {
	// read file a.txt
	data, err := os.ReadFile("a.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(data))

	// create file b.txt
	err = os.WriteFile("b.txt", data, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
}
