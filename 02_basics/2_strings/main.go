package main

import "fmt"

func main() {
	// Interpreted string
	a := "\nCiao Mariarosa,\n\tcome stai?\n\nMario"
	fmt.Println(a)

	// Raw string
	a = `Say "Hello"`
	fmt.Println(a)

	// Multiline raw string
	a = `# json string
	{
		"key": "value"
	}`
	fmt.Println(a)

	var b rune = 'a'
	fmt.Printf("%v %c\n", b, b)
}
