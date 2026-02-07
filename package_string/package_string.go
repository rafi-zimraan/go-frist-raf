package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello World", "World"))
	fmt.Println(strings.ToLower("Hello World"))
	fmt.Println(strings.ToUpper("Hello World"))
	fmt.Println(strings.Trim(".  Hello World.       ", " ."))
	fmt.Println(strings.ReplaceAll("Hello World", "World", "Rafi"))
	fmt.Println(strings.Replace("Hello World", "World", "Rafi", 1))
}
