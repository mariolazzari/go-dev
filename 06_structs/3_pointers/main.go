package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) String() {
	fmt.Println(u.Name, u.Age)
}

// arg by value
func changeName(u User, name string) {
	u.Name = name
	fmt.Println("Name inside changeName:", u.Name)
}

// arg by reference
func changeNameRef(u *User, name string) {
	u.Name = name
	fmt.Println("Name inside changeName:", u.Name)
}

func main() {
	var mario User = User{"Mario", 50}
	mario.String()

	changeName(mario, "Marco")
	mario.String()

	changeNameRef(&mario, "Marco")
	mario.String()
}
