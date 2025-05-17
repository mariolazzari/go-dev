package main

import (
	"fmt"
	"sync"
	"time"
)

func generateThumb(image string, size int) {
	thumb := fmt.Sprintf("%s@%dx.png", image, size)
	fmt.Println("Generate thumb")
	time.Sleep(time.Microsecond * time.Duration(size))
	fmt.Println("Generated thumb", thumb)
}

func main() {
	const img = "foo.png"
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(i)
		go func(i int) {
			defer wg.Done()
			generateThumb(img, i+1)
		}(i)
	}

	fmt.Println("Witing for thumbs...")
	wg.Wait()
	fmt.Println("thumbs done")
}
