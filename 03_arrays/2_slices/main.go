package main

import "fmt"

func main() {
	names := []string{}
	fmt.Println("Capacity:", cap(names))
	fmt.Println("Length:", len(names))

	names = append(names, "John")
	names = append(names, "Jane")
	fmt.Println("\nCapacity:", cap(names))
	fmt.Println("Length:", len(names))

	names = append(names, "Doe")
	fmt.Println("\nCapacity:", cap(names))
	fmt.Println("Length:", len(names))

	names = append(names, "Mario")
	names = append(names, "Luigi")
	fmt.Println("\nCapacity:", cap(names))
	fmt.Println("Length:", len(names))

	// make defines a slice with a length and capacity
	names = make([]string, 5, 10)
	fmt.Println("\nCapacity:", cap(names))
	fmt.Println("Length:", len(names))

	// subset of a slice
	subset := names[1:3]
	fmt.Println("\nSubset:", subset)
	fmt.Println("Subset Capacity:", cap(subset))
	fmt.Println("Subset Length:", len(subset))
	subset[0] = "Peach"

	// subset are references to the original slice
	fmt.Println("\nSubset:", subset)
	fmt.Print("Names:", names)

	// copy: clone a slice
	namesCopy := make([]string, len(names))
	copy(namesCopy, names)
	fmt.Println("\n\nNames Copy:", namesCopy)
	namesCopy[0] = "Yoshi"
	fmt.Println("\nNames Copy:", namesCopy)
	fmt.Print("Names:", names)
}
