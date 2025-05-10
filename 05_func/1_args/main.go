package main

import "fmt"

func main() {
	hello("Mario")
	greeting("Hi", "Mario")

	hello := sayHello()
	fmt.Println(hello)

	names := []string{"Mario", "Mariarosa", "Maria"}
	gs, len, cap := info(names)
	fmt.Println(gs)
	fmt.Println(len)
	fmt.Println(cap)

	isValid()

}

// single arg
func hello(name string) {
	fmt.Println("Hello", name)
}

// multiple args
func greeting(greet, name string) {
	fmt.Printf("%s %s\n", greet, name)
}

// return val
func sayHello() string {
	return "Hello"
}

// mutl return vals
func info(s []string) (string, int, int) {
	gs := fmt.Sprintf("%#v", s)
	l := len(s)
	c := cap(s)

	return gs, l, c
}

// named return
func isValid() (valid bool) {
	fmt.Println(valid)
	valid = true
	fmt.Println(valid)

	return
}
