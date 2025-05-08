package main

import "fmt"

func main() {
	namesArrray := [4]string{"John", "Jane", "Doe", "Mario"}
	fmt.Println(namesArrray)

	nameSlice := []string{"John", "Jane", "Doe", "Mario"}
	fmt.Println(nameSlice)

	// append to slice
	nameSlice = append(nameSlice, "Luigi")
	fmt.Println(nameSlice)
	nameSlice = append(nameSlice, "Peach", "Toad")
	fmt.Println(nameSlice)
	nameSlice[0] = "Yoshi"
	fmt.Println(nameSlice)

	a1 := []int{1, 2, 3, 4, 5}
	var a2 []int

	a2 = a1
	a1[0] = 0
	fmt.Println(a1) // [0 2 3 4 5]
	fmt.Println(a2) // [10 2 3 4 5]

	// variant operator
	a3 := []int{1, 2, 3}
	a4 := []int{4, 5, 6}
	a5 := append(a3, a4...)
	fmt.Println(a5) // [1 2 3 4 5 6]
}
