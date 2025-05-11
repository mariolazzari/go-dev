package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) String() {
	fmt.Println(u.Name, u.Age)
}

func main() {
	var mario User = User{"Mario", 50}
	mario.String()
}
