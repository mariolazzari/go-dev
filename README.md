# Go for developers

## Packages, modules and dependencies

### Modules

A Go module is a collection of Go packages and their dependencies that can be built, versioned and managed as a unit.

```sh
go mod init demo
go mod init github.com/user/repo
``` 

### Packages

A package is a collection of go files that share the same import path. This files are *always* in the same directory and this directory is *always* the name of the package.

- lower cased
- letters and numbers only
- no starting numebrs
- declared at the top of file

### Files and directories

- singularly focused
- lowecase and _
- clean and simple to use

### Dependency management

```sh
go mod init demo
go get github.com/gobuffalo/flect
```

Go sum file will be generated: ```go go get -u``` will update module.

## Basics

### Numbers

- Signed
- Unsigned

### Strings

- Interpreted strings: ""
- Raw strings: ``
- UTF-8 support

```go
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
}
```

### Variables

- var name type
- :=

### Constants

- const

### Naming

- case sensitive
- conventional
- very short
- Uppercase: public

### Printing & formatting

- fmt package [docs](https://pkg.go.dev/fmt)

## Arrays, slices and iterations

### Arrays and slices

|            | Arrays | Slices |
| ---------- | ------ | ------ |
| Fixed size | Yes    | No     |
| Fixed type | Yes    | Yes    |
| Zero based | Yes    | No     |

### How slices work

- len [docs](https://pkg.go.dev/builtin#len)
- cap [docs](https://pkg.go.dev/builtin#cap)
- make [docs](https://pkg.go.dev/builtin#make)
- copy [docs](https://pkg.go.dev/builtin#copy)

```go
type Slice struct {
	// number of actual elements
	Length int
	// maximum number of elements
	Capacity int
	// actual array
	Array []string
}
```

### Iteration

- for
- range

```go
    for i := 0; i < 10; i++ {
		fmt.Println(i)
    }

	for i := range 5 {
		fmt.Println(i)
	}

    names := [5]string{"John", "Jane", "Doe", "Smith", "Alice"}
	for i, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", i, name)
	}
```

## Maps and contrul structures

### Maps

Unordered set on key/value pairs.

```go
    users := map[string]string{}
    users["Mario"] = "mario.lazzari@gmail.com"

    // init
	users2 := map[string]string{
		"Mario": "mario",
		"Luigi": "luigi",
		"Peach": "peach",
	}

    // make
	users3 := make(map[string]string)
	users3["Mario"] = "mario"
	users3["Luigi"] = "luigi"
	users3["Peach"] = "peach"

    // order is not guaranteed
	for k, v := range users {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

    // delete
	delete(users, "Mario")
	fmt.Println(users)

    // extract keys
    months :

    // extract keys
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	keys := make([]int, 0, len(months))
	for k := range months {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

```

### If

```go
    greet := true
	if greet {
		println("Hello, World!")
	} else {
		println("Goodbye, World!")
	}
```

