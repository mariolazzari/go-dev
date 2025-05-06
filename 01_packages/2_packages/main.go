package main

import (
	"fmt"
	"path"
)

func main() {
	a := path.Join("a", "b")
	fmt.Println(a)

	b := path.Join("a", "b")
	fmt.Println(b)
}
