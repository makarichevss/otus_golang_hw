package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	var greeting = "Hello, OTUS!"
	fmt.Println(reverse.String(greeting))
}
