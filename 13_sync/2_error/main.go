package main

import (
	"fmt"
	"sync"
	"time"
)

func generateThumb(image string, size int) error {
	if size == 0 {
		return fmt.Errorf("Invalid size %d", size)
	}

	thumb := fmt.Sprintf("%s@%dx.png", image, size)
	fmt.Println("Generate thumb")
	time.Sleep(time.Microsecond * time.Duration(size))
	fmt.Println("Generated thumb", thumb)
	return nil
}

func main() {
	const img = "foo.png"
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if err := generateThumb(img, i+1); err != nil {
				fmt.Println("Error:", err)
			}
		}(i)
	}

	fmt.Println("Witing for thumbs...")
	wg.Wait()
	fmt.Println("thumbs done")
}
