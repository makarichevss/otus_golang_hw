package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse" //nolint:gomnd
)

func main() {
	const greeting = "Hello, OTUS!"
	fmt.Println(reverse.String(greeting))
}
