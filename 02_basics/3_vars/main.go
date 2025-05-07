package main

import "fmt"

func main() {
	var i int
	var f float64
	var s string
	var b bool

	// Print the zero values
	fmt.Println(i) // 0
	fmt.Println(f) // 0
	fmt.Println(s) // ""
	fmt.Println(b) // false

	i = 42
	f = 3.14
	s = "Hello, World!"
	b = true

	// Print the values
	fmt.Println(i) // 42
	fmt.Println(f) // 3.14
	fmt.Println(s) // "Hello, World!"
	fmt.Println(b) // true

	// Short variable declaration
	x := 42
	y := 3.14
	z := "Hello, World!"
	w := true
	// Print the values
	fmt.Println(x) // 42
	fmt.Println(y) // 3.14
	fmt.Println(z) // "Hello, World!"
	fmt.Println(w) // true

	// Type conversion
	y = float64(x)
	fmt.Println(i) // 42
}
