package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer fmt.Println("Mario")
	fmt.Println("Ciao")

	creatFile("mario.txt")

	defer func(now time.Time) {
		fmt.Printf("duration: %s\n", time.Since(now))
	}(time.Now())

	fmt.Println("Sleep for 50ms")
	time.Sleep(50 * time.Millisecond)
}

func creatFile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString("Ciao Mario")
	if err != nil {
		return err
	}

	return nil

}
