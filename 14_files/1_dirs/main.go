package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir("data")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("->", file.Name())
			continue
		}
		fmt.Println(file.Name())
		info, err := file.Info()
		if err != nil {
			log.Println("Error getting file info:", err)
			continue
		}
		fmt.Printf("%+v\n", info)
	}
}
