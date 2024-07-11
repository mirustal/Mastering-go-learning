package main

import (
	"fmt"
	"strings"
)

func main() {
	f := fmt.Printf

	f("equal: %v\n", strings.EqualFold("hello", "hello"))
	f("equal: %v\n", strings.EqualFold("helol", "hellO"))
	f("prefix: %v\n", strings.HasPrefix("helol", "He"))
	f("sufix: %v\n", strings.HasSuffix("helol", "lo"))

	f(strings.Fields("hello wrold")[])

}