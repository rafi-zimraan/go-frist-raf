package main

import "fmt"

func main() {
	type NIK string

	var NikBudi NIK = "3275027008190004"

	var NikAndi string = "3275027008190005"
	var ContohNik NIK = NIK(NikAndi)

	fmt.Println(NikBudi)
	fmt.Println(ContohNik)
}
