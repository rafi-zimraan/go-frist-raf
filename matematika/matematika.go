package main

import "fmt"

func main() {

	var a int = 20
	var b int = 30
	var c int = 10
	var d int = 5
	var e = a - b - c*d
	fmt.Println(e)

	var i = 20
	i += 30
	fmt.Println(i)

	i *= 2
	fmt.Println(i)

}
