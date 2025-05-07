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