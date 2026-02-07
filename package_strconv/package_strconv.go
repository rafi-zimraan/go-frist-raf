package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.Atoi("salah")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(result)
	}

	resultBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(resultBool)
	}
}
