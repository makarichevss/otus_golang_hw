package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	const greeting = "Hello world"
	fmt.Println(reverse.String(greeting))
}
