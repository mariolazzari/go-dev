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

### Switch

```go
    switch month {
	case 1:
		println("January")
	case 2:
		println("February")
	case 3:
		println("March")
	case 4:
		println("April")
	case 5:
		println("May")
	case 6:
		println("June")
	case 7:
		println("July")
	case 8:
		println("August")
	case 9:
		println("September")
	case 10:
		println("October")
	case 11:
		println("November")
	case 12:
		println("December")
	default:
		println("Invalid month")
	}
```

## Functions

### Arguments and return values

```go
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
```

### First class functions

```go
package main

import "fmt"

func main() {
	fn := func() string {
		return "Hello"
	}

	hello := fn()
	fmt.Println(hello)

	sayHello(fn)

	name := "Mario"
	hi := func() {
		fmt.Println("Hi", name)
	}
	hi()
}

func sayHello(f func() string) {
	fmt.Println(f())
}
```

### Variadic arguments

```go
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
```

### Deferring function calls

```go
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
```

### Init function

```go
package main

import "fmt"

func main() {
	fmt.Println("Main")
}

// init package
func init() {
	fmt.Println("Init")
}
```

## Structs, methods and pointers

### Structs

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type MyMap map[string]string

type User struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	var myMap MyMap = MyMap{"mario": "Mario"}
	fmt.Println(myMap)

	var user User = User{"Mario", 50}
	fmt.Println(user)

	var maria User = User{
		Name: "Maria",
		Age:  78,
	}
	fmt.Println(maria)

	enc := json.NewEncoder(os.Stdout)

	if err := enc.Encode(maria); err != nil {
		log.Fatal(err)
	}

}
```

### Methods

```go
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
```

### Pointers

```go
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
```

## Testing

### Basics

```sh go test```

Go [test](https://pkg.go.dev/testing) must:
- be in a *_test.go* file.
- start with prefix *Test*
- accept *Testing.T* type
- return no args

### Running tests

```sh
# run all tests
go test ./...
# run one package tests
go test ./package_name
# verbose
go test -v ./... 
# run in parallel
go test -v -run=regex
# run with timeout
go test -v -timeout=50ms
```

### Coverage

```sh
go test -cover
# output
go test -coverprofile=coverage.out
# format output
go tool cover -func=coverage.out
# html
go tool cover -html=coverage.out
```

### Table driven

### Helpers

## Interfaces

### What are inrerfaces?

Collection of methods (not fields):
- concrete type does not need to know about interfaces
- you can write interface for existing types
- you can write interface for third party types

### Empty interface

```go
package main

// generic interface
var i interface{}
var a any

// named empty interface
type foo interface{}

func main() {

}
```

### Using interfaces

```go
type Model interface {
	ID() int
}

type Data map[int]Model

type Store struct {
	data Data
}

func (s *Store) Insert(m Model) error {
	s.data[m.ID()] = m
	return nil
}```

IO [package](https://pkg.go.dev/io) includes many interfaces.

### Type assertions

```go
func WriteNow(i any, s string) error {
	w, ok := i.(io.Writer)
	if !ok {
		return fmt.Errorf("i is not a io.Writer")
	}

	_, err := fmt.Fprintln(w, s)
	if err != nil {
		return err
	}

	return nil
}
```

## Errors

### Basics

Go treats errors as values.

```go
func Get(key string) (string, error) {
	m := map[string]string{
		"a": "A",
		"b": "B",
	}

	if v, ok := m[key]; ok {
		return v, nil
	}

	return "", fmt.Errorf("Key not found: %s", key)
}
```

### Panics

Function with a string argument.

- panic crash application
- recover from panic ang log error
- capture panic and return error

- use defer with a recover to catch panic
- use type assertion for error checking
- use named return
- never raise a panic

```go
package main

func DoSomething(input int) (err error) {

	defer func() {
		p := recover()
		if p == nil {
			return
		}

		if e, ok := p.(error); ok {
			err = e
			return
		}

		//		err := fmt.Errorf("panit %T %s", p, p)
	}()

	switch input {
	case 0:
		return nil
	case 1:
		panic("one")
	}

	return nil
}

func main() {

}
```

### Custom errors

```go
type ErrTableNotFound struct {
	Table      string
	OccurredAt time.Time
}

func (e ErrTableNotFound) Error() string {
	return fmt.Sprintf("%s table not found %s", e.OccurredAt, e.Table)
}
```

### Wrapping errors

- [Wrap](https://pkg.go.dev/github.com/pkg/errors#Wrap)
- [Unwrap](https://pkg.go.dev/github.com/pkg/errors#Unwrap)


### Using As and Is

- As [method](https://pkg.go.dev/errors#As)
- Is [method](https://pkg.go.dev/errors#Is)

## Generics

### What is a generic

```go
func keys(m map[any]any) []any {
	keys := make([]any, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
```

### Type constraints

```go
func Slicer[T any](input []T) []T {
	return input
}
```

### Defining constraints

```go
type MapKey interface {
	int | float64
}

func Keys[K MapKey, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
```

### Underlying type constraints

*~* operator

```go
type MyInit int

type MapKey interface {
	~int | float64
}

func Keys[K MapKey, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
```

### Generic types

```go
import "golang.org/x/exp/constraints"

type Store[K constraints.Ordered] struct {
	items []K
}
```

## Channels

### Concurrency

- Concurrency is the composition of indipendetly execution tasks: it is about structure.
- Parallelism is doing multiple task at once.

```go
package main

import "fmt"

func main() {

	go someFunc()
	go func() {
		fmt.Println("Anonymous func")
	}()

}

func someFunc() {
	fmt.Println("Some func")
}
```

### Channels

- typed
- synchronous transmissions
- fifo
- unbuffered or buffered
- directional

```go
package main

import "fmt"

func main() {
	phone := make(chan string)
	defer close(phone)

	go Janis(phone)
	phone <- "Hello Janis"
	msg := <-phone
	fmt.Println("Janis said:", msg)
}

func Janis(ch chan string) {
	msg := <-ch
	fmt.Println("Jimi said:", msg)
	ch <- "Hello Jimi"
}
```

### Iteration and select 

```go
package main

import "fmt"

func main() {
	const N = 5

	ch := make(chan int)

	for i := range N {
		go func(i int) {
			for m := range ch {
				fmt.Printf("Routine %d received %d\n", i, m)
			}
		}(i)
	}

	close(ch)
}
```

### Closed channel

```go
package main

import (
	"fmt"
	"time"
)

func listener(i int, quit <-chan struct{}) {
	fmt.Printf("listener %d is waiting\n", i)
	<-quit
	fmt.Printf("listener %d is exiting\n", i)
}

func main() {
	quit := make(chan struct{})

	for i := range 5 {
		go listener(i, quit)
	}

	time.Sleep(10 * time.Microsecond)
	fmt.Println("Closing quit")
	close(quit)
	time.Sleep(50 * time.Microsecond)

}
```

### Buffered channels

Message left by caller that it may be picked up later by the receiver.

```go
package main

import "fmt"

func main() {
	message := make(chan string, 2)
	message <- "hello"
	message <- "hello again"
	fmt.Println(<-message)
	fmt.Println(<-message)
}
```

### System signals

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	fmt.Println("Awaiting signal...")
	s := <-ch
	fmt.Println("Got signal", s)
}
```

## Context

### Contexts

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Printf("%v\n", ctx)
	fmt.Printf("\t%#v\n", ctx)
}
```

### Context values

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	fmt.Println(ctx.Value("key"))
}
```