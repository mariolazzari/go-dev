package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {

		if i == 3 {
			continue
		}

		if i == 8 {
			break
		}

		fmt.Println(i)
	}

	for i := range 5 {
		if i == 4 {
			break
		}

		fmt.Println(i)
	}

	names := [5]string{"John", "Jane", "Doe", "Smith", "Alice"}
	for i, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", i, name)
	}

	// just index
	for i := range names {
		fmt.Printf("Index: %d\n", i)
	}
	// just value
	for _, name := range names {
		fmt.Printf("Name: %s\n", name)
	}

}
