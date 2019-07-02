package main
// Package == project == workspace. first line must always contain its package at the top
// Two types of packages: executable--generates an executable, reusable--'helpers', or reusable logic
// go build only works here because of package main. It instructs go to compile some usable code

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

// How to run? Some of the commands available with the GO CLI
// go build: compile source code
// go run: compile and execute one or two files
// go fmt: format all code in the current dir
// go install: compile and install a package
// go get: download raw source code of someone else's package
// go test: run any tests associated with the current project