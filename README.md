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

