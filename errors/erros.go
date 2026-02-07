package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("validation error")
	notFoundError   = errors.New("not found error")
)

func GetByFood(food string) error {
	if food == "" {
		return validationError
	}

	if food != "nasi goreng" {
		return notFoundError
	}

	// sukses
	return nil
}

func main() {
	err := GetByFood("nasi goreng")

	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, notFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}
