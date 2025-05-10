package main

import "fmt"

func main() {
	names := []string{"Mario", "Mariarosa", "Maria"}
	hello(names...)

	id1 := 1
	id2 := 2
	id3 := 3
	ids := []int{id1, id2, id3}

	loadUsers(id1)
	loadUsers(id1, id2, id3)
	loadUsers(ids...)
}

// variadic argument must be last one
func hello(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func loadUsers(ids ...int) {
	fmt.Println("looking up:", ids)
}
