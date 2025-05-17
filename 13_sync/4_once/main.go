package main

import (
	"fmt"
	"sync"
	"time"
)

type Builder struct {
	Built bool
	once  sync.Once
}

func (b *Builder) Build() error {
	b.once.Do(func() {
		fmt.Println("Building...")
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Built")
		b.Built = true
	})

	return nil
}

func main() {

}
